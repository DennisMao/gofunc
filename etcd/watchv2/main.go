package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/client"
)

var (
	kapi client.KeysAPI
	c    client.Client
)

func Init() error {
	cfg := client.Config{
		Endpoints: []string{"http://192.168.33.10:32379", "http://192.168.33.10:12379", "http://192.168.33.10:22379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		panic(err)
	}
	kapi = client.NewKeysAPI(c)
	return nil
}

func Set(k, v string) error {

	resp, err := kapi.Set(context.Background(), k, v, nil)
	if err != nil {
		return err
	}

	return nil
}

func Get(k string) string {
	resp, err := kapi.Get(context.Background(), k, nil)
	if err != nil {
		log.Printf("[Get] error:%s \n", err.Error())
		return ""
	}

	return resp.Node.Value
}

// 监听
func Watch(k string, f func(act, k, v string, modIndex uint64)) error {
	w := kapi.Watcher(k, nil)
	var err error = nil

	for err == nil {
		resp, err := w.Next(context.Background())
		if err != nil {
			return err
		}

		f(resp.Action, resp.Node.Key, resp.Node.Value, resp.Node.ModifiedIndex)
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
	go Watch(key, func(act, k, v string, modIndex uint64) {
		log.Printf("[Watch] act:%s key:%s value:%s modIndex:%d \n", act, key, value, modIndex)
	})

	log.Println("Get key:%s , value:%s exceptValue:%s", key, Get(key), value)

	Set(key, value)

	log.Printf("Set key:%s , value:%s \n", key, value)
	log.Printf("Get key:%s , value:%s exceptValue:%s \n", key, Get(key), value)
}
