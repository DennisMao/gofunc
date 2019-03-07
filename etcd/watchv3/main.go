package main

import (
	"context"
	"log"
	"time"

	client "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

var (
	c *client.Client
)

func Init() error {
	cfg := client.Config{
		Endpoints: []string{"http://192.168.33.10:32379", "http://192.168.33.10:12379", "http://192.168.33.10:22379"},
	}

	cli, err := client.New(cfg)
	if err != nil {
		panic(err)
	}

	c = cli
	return nil
}

func Set(k, v string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err := c.Put(ctx, k, v)
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}

	return nil
}

func Get(k string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// 支持多种请求设置:
	// 1.前缀  	clientv3.WithPrefix()
	// 2.排序  	clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend)
	// 3.限制长度	clientv3.WithLimit()
	// 4.统计	clientv3.WithCountOnly()
	//
	// Mod version相关
	// 1.指定rev clientv3.WithRev()
	// 2.最大rev clientv3.WithMaxModRev()
	// 3.最小rev clientv3.WithMaxMinRev()

	resp, err := c.Get(ctx, k)
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}

	// 与v2版本区别比较大的是,v3版本的响应中支持多个键值对返回,因此在第一层级的响应有Count
	// 真正内容在数组`Kvs []*mvccpb.KeyValue`中。
	if resp.Count == 0 {
		return ""
	}

	return string(resp.Kvs[0].Value)
}

// 监听
func Watch(k string, f func(act, k, v string, modIndex int64)) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	w := c.Watch(ctx, k)

	for wresp := range w {

		if wresp.Err() != nil {
			return wresp.Err()
		}

		for _, ev := range wresp.Events {
			// Call back
			f("put", string(ev.Kv.Key), string(ev.Kv.Value), ev.Kv.Version)
		}

	}

	return nil
}

func main() {
	//Initial etcd client
	err := Init()
	if err != nil {
		panic(err)
	}

	key := "/testKey"
	value := "1"

	// 开启监听协程
	go Watch(key, func(act, k, v string, modIndex int64) {
		log.Printf("[Watch] act:%s key:%s value:%s modIndex:%d \n", act, key, value, modIndex)
	})

	log.Println("Get key:%s , value:%s exceptValue:%s", key, Get(key), value)

	Set(key, value)

	log.Printf("Set key:%s , value:%s \n", key, value)
	log.Printf("Get key:%s , value:%s exceptValue:%s \n", key, Get(key), value)

}
