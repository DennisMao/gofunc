// Redis集群客户端 Radix
//
// 经验:
//     1.由于radix包内部对每个连接池都会开单独协程进行维护,需要合理控制连接池的数量和大小
//     2.如果集群状态稳定,或者业务对一致性要求没那么高的,建议降低集群槽数据同步频率,减少性能开销
//     3.Redis若对客户端本身没做timeout处理的,在连接池内可不单独设置PingPong,减少性能开销
//     4.在集群需要进行rebalance或者reshard操作时候,建议从业务层面发送空指令触发Radix客户端集群的槽数据同步,减少对业务延迟的影响
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mediocregopher/radix"
)

func InitCluster(addrs string, poolSize, pipelineLimit int, auth string) (*radix.Cluster, error) {

	peers := strings.Split(addrs, ",")

	fmt.Printf("Start to init cluster on endPoints:%s", addrs)
	clu, err := radix.NewCluster(peers, radix.ClusterPoolFunc(radix.ClientFunc(func(network, addr string) (radix.Client, error) {

		poolSize := poolSize            // 池大小
		pingInterval := 5 * time.Second // 心跳探测周期

		opts := []radix.PoolOpt{}
		opts = append(opts, radix.PoolPingInterval(pingInterval))
		opts = append(opts, radix.PoolPipelineConcurrency(pipelineLimit))
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

	return clu, nil
}

func main() {

	addrs := "127.0.0.1:30001,127.0.0.1:30002,127.0.0.1:30003"
	clients := 10  // 连接池客户端数量
	pipeline := 10 // 连接池并发限制
	auth := ""     // Redis密码

	cli, err := InitCluster(addrs, clients, pipeline, auth)
	if err != nil {
		panic(err)
	}

	var ret string
	err = cli.Do(radix.Cmd(&ret, "SET", "1", "1"))
	if err != nil {
		fmt.Printf("执行:SET 1 1 结果:%s \n", err.Error())
		return
	}

	fmt.Printf("执行:SET 1 1 结果:%s \n", ret)

}
