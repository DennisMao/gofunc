// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"strings"

	"raft-example/app/http"
	"raft-example/app/models/kvstore"
	"raft-example/app/models/raft"

	"go.etcd.io/etcd/raft/raftpb"
)

func main() {
	// 配置项
	cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers") //集群地址,例:http://ip:123,http://ip:234
	id := flag.Int("id", 1, "node ID")                                                          //当前节点编号
	kvport := flag.Int("port", 9121, "key-value server port")                                   //Kv数据库服务端口,API服务端口
	join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()

	proposeC := make(chan string)
	defer close(proposeC)
	confChangeC := make(chan raftpb.ConfChange)
	defer close(confChangeC)

	// raft provides a commit stream for the proposals from the http api
	var kvs *kvstore.KvStore
	getSnapshot := func() ([]byte, error) { return kvs.GetSnapshot() }
	commitC, errorC, snapshotterReady := raft.NewRaftNode(*id, strings.Split(*cluster, ","), *join, getSnapshot, proposeC, confChangeC)

	kvs = kvstore.NewKvStore(<-snapshotterReady, proposeC, commitC, errorC)

	// the key-value http handler will propose updates to raft
	http.ServeHttpKVAPI(kvs, *kvport, confChangeC, errorC)
}
