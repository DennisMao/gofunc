// Redis集群压测工具
// 使用方法
//         1. go mod tidy
// 		   2. go build
//         3. radix --help //查看使用帮助,参数按照redis-benchmark实现
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mediocregopher/radix"
)

var (
	gClu   Client
	config Config
)

type Config struct {
	ClusterAddrs string // 集群节点地址
	Clients      int    // 连接池客户端数量
	Requests     int    //请求总量
	Datasize     int    // 写入值得长度 bytes
	Pipeline     int    // 并发写入
	Keyspacelen  int    // 随机生成key的长度
	Auth         string // 集群密码
}

type Client interface {
	Do(a radix.Action) error
}

func init() {
	defaultPath := "127.0.0.1:30001,127.0.0.1:30002,127.0.0.1:30003"

	flag.StringVar(&config.ClusterAddrs, "h", defaultPath, "all nodes ips,>=3 use cluster")
	envClusterPeers := os.Getenv("REDIS_CLUSTER")
	if envClusterPeers != "" {
		config.ClusterAddrs = envClusterPeers
	}

	flag.IntVar(&config.Clients, "c", 50, "Number of parallel connections (default 50)")
	flag.IntVar(&config.Requests, "n", 100000, "Total number of requests (default 100000)")
	flag.IntVar(&config.Datasize, "d", 3, "Data size of SET/GET value in bytes (default 3)")
	flag.IntVar(&config.Keyspacelen, "r", 12, `Use random keys for SET/GET/INCR, random values for SADD
  Using this option the benchmark will expand the string __rand_int__
  inside an argument with a 12 digits number in the specified range
  from 0 to keyspacelen-1. The substitution changes every time a command
  is executed. Default tests use this to hit random keys in the
  specified range.`)
	flag.IntVar(&config.Pipeline, "P", 1, "Pipeline <numreq> requests. Default 1 (no pipeline).")
	flag.StringVar(&config.Auth, "a", "", `Password Password to use when connecting to the server.
  You can also use the REDISCLI_AUTH environment
  variable to pass this password more safely
  (if both are used, this argument takes predecence).`)

}

func InitCluster(addrs string, poolSize, pipelineLimit int, auth string) (Client, error) {

	peers := strings.Split(addrs, ",")

	if len(peers) == 1 {
		clu, err := radix.NewPool("tcp",
			peers[0],
			poolSize,
			radix.PoolPingInterval(time.Second*5),
			radix.PoolPipelineConcurrency(pipelineLimit))
		return clu, err
	} else {
		fmt.Printf("Start to init cluster on endPoints:%s", config.ClusterAddrs)
		clu, err := radix.NewCluster(peers, radix.ClusterPoolFunc(radix.ClientFunc(func(network, addr string) (radix.Client, error) {
			//  客户端连接池管理
			//	PoolConnFunc(DefaultConnFunc)
			//	PoolOnEmptyCreateAfter(1 * time.Second)
			//	PoolRefillInterval(1 * time.Second)
			//	PoolOnFullBuffer((size / 3)+1, 1 * time.Second)
			//	PoolPingInterval(5 * time.Second / (size+1))
			//	PoolPipelineConcurrency(size)
			//	PoolPipelineWindow(150 * time.Microsecond, 0)
			//
			poolSize := poolSize                 // 池大小
			pingInterval := 5 * time.Second      // 心跳探测周期
			pipelineConcurrency := pipelineLimit // pipeline的并发数

			opts := []radix.PoolOpt{}
			opts = append(opts, radix.PoolPingInterval(pingInterval))
			opts = append(opts, radix.PoolPipelineConcurrency(pipelineConcurrency))
			if auth != "" {
				opts = append(opts, radix.PoolConnFunc(func(network, addr string) (radix.Conn, error) {
					return radix.Dial(network, addr, radix.DialAuthPass(auth))
				}))
			}

			p, err := radix.NewPool(
				network,
				addr,
				poolSize,
				opts...,
			)
			if err != nil {
				fmt.Printf("radix: create pool for client addr:%s poolsize:%d failed:%s \n", addr, poolSize, err)
				return nil, err

			}
			fmt.Printf("radix: create pool for client addr:%s poolsize:%d Success!!\n", addr, poolSize)

			return p, nil
		})))
		if err != nil {
			return nil, err
		}

		go func() {
			for {
				select {
				case err := <-clu.ErrCh:
					fmt.Println(err)
					os.Exit(-1)
				}
			}
		}()

		return clu, nil
	}

}

func main() {

	flag.Parse()

	cli, err := InitCluster(config.ClusterAddrs, config.Clients, config.Pipeline, config.Auth)
	if err != nil {
		panic(err)
	}

	configJson, _ := json.MarshalIndent(config, "", " ")
	fmt.Printf("%s \n", string(configJson))

	// 生成随机长度keys/values
	values := generateRandomValue(config.Datasize)
	keys := generateRandomKeys(config.Keyspacelen, config.Requests)
	pipelineToken := make(chan struct{}, config.Pipeline) // 漏桶 控制并发数量
	for i := 0; i < config.Pipeline; i++ {
		pipelineToken <- struct{}{}
	}
	wg := sync.WaitGroup{}
	totalFinish := int64(0)
	ctx, finsih := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Currenct Finish :%d / %d \n", totalFinish, config.Requests)
				return
			default:
			}

			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Currenct Finish :%d / %d \n", totalFinish, config.Requests)
		}
	}(ctx)

	tS := time.Now() //计时器-开始
	for i, _ := range keys {

		<-pipelineToken
		wg.Add(1)
		go func(key string) {
			defer func() {
				pipelineToken <- struct{}{}
				wg.Done()
			}()

			err := cli.Do(radix.Cmd(nil, "SET", key, values))
			if err != nil {
				fmt.Printf("执行失败 SET:%s :%s,详情:%s \n", key, values, err.Error())
				return
			} else {
				atomic.AddInt64(&totalFinish, 1)
			}

		}(keys[i])
	}

	wg.Wait()
	finsih()
	costTime := time.Now().Sub(tS).Nanoseconds() // 计时器结束

	// 生成结果  按照redis benchmark
	fmt.Printf("====== %s ======\n", "BenchmarkReport")
	fmt.Printf("  %d requests completed in %.4f seconds\n", totalFinish, float64(costTime)/1000000000.0)
	fmt.Printf("  %d parallel clients\n", config.Clients)
	fmt.Printf("  %d parallel pipeline\n", config.Pipeline)
	fmt.Printf("  %d bytes payload\n", config.Datasize)
	fmt.Printf("  %f requests per second\n", (float64(totalFinish)*1000000000.0)/float64(costTime))

}

// 生成随机values
// 生生成随机key
func generateRandomValue(size int) string {
	val := make([]byte, size)
	for i := 0; i < size; i++ {
		val[i] = 'x'
	}
	return string(val)
}

// 生生成随机key
func generateRandomKeys(keyspacelen, requests int) []string {
	keysChan := make([]string, requests)

	for i := 0; i < requests; i++ {
		keysChan[i] = fmt.Sprintf("key:%012d", i)
	}

	return keysChan
}
