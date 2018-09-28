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

package kvstore

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"sync"

	"go.etcd.io/etcd/etcdserver/api/snap"
)

// a key-value store backed by raft
type KvStore struct {
	proposeC    chan<- string // channel for proposing updates
	mu          sync.RWMutex
	KvStore     map[string]string // current committed key-value pairs
	snapshotter *snap.Snapshotter
}

type kv struct {
	Key string
	Val string
}

// 实例化KV存储
//
// @snapshotter: 快照器,用于本地日志快照和恢复
// @commitC: 提交信息通道
// @errorC: 错误信息返回通道
func NewKvStore(snapshotter *snap.Snapshotter, proposeC chan<- string, commitC <-chan *string, errorC <-chan error) *KvStore {
	s := &KvStore{proposeC: proposeC, KvStore: make(map[string]string), snapshotter: snapshotter}
	// replay log into key-value map
	s.readCommits(commitC, errorC)
	// read commits from raft into KvStore map until error
	go s.readCommits(commitC, errorC)
	return s
}

// 查询数据
func (s *KvStore) Lookup(key string) (string, bool) {
	s.mu.RLock()
	v, ok := s.KvStore[key]
	s.mu.RUnlock()
	return v, ok
}

// 将待返回数据打包进入预备队列
func (s *KvStore) Propose(k string, v string) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(kv{k, v}); err != nil {
		log.Fatal(err)
	}
	s.proposeC <- buf.String()
}

// 读取提交请求
//
// @commitC: 提交信息通道,内容由gob编码,格式为 {Key:string,Val:string}
// @errorC:  错误返回信息通道
func (s *KvStore) readCommits(commitC <-chan *string, errorC <-chan error) {
	for data := range commitC {
		if data == nil {
			// done replaying log; new data incoming
			// OR signaled to load snapshot
			snapshot, err := s.snapshotter.Load()
			if err == snap.ErrNoSnapshot {
				return
			}
			if err != nil {
				log.Panic(err)
			}
			log.Printf("loading snapshot at term %d and index %d", snapshot.Metadata.Term, snapshot.Metadata.Index)
			if err := s.RecoverFromSnapshot(snapshot.Data); err != nil {
				log.Panic(err)
			}
			continue
		}

		var dataKv kv
		dec := gob.NewDecoder(bytes.NewBufferString(*data))
		if err := dec.Decode(&dataKv); err != nil {
			log.Fatalf("raftexample: could not decode message (%v)", err)
		}
		s.mu.Lock()
		s.KvStore[dataKv.Key] = dataKv.Val
		s.mu.Unlock()
	}
	if err, ok := <-errorC; ok {
		log.Fatal(err)
	}
}

// 获取快照
// 将当前kv内容直接json序列化一份返回
func (s *KvStore) GetSnapshot() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return json.Marshal(s.KvStore)
}

// 从快照恢复
// 将原数据json反序列化后直接替换掉当前kv内存指针
func (s *KvStore) RecoverFromSnapshot(snapshot []byte) error {
	var store map[string]string
	if err := json.Unmarshal(snapshot, &store); err != nil {
		return err
	}
	s.mu.Lock()
	s.KvStore = store
	s.mu.Unlock()
	return nil
}
