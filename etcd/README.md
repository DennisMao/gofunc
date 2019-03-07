# Etcd

## 部署

### 单机集群

使用goreman同时启动三个进程,为别为etcd3个节点。


step1 安装etcd和goreman
1. etcd : https://github.com/etcd-io/etcd/releases
2. goreman: https://github.com/mattn/goreman/releases

step2 创建目录用于存放etcd产生的日志和wal
```
mkdir -p node1/data node2/data node3/data
```


step3 新建名为`Procfile`文件,启动脚本,脚本内容如下
```
node1: etcd -name etcd1 --data-dir $PWD/node1/data --listen-client-urls http://0.0.0.0:12379 --listen-peer-urls http://0.0.0.0:12380 --advertise-client-urls http://0.0.0.0:12379 --initial-advertise-peer-urls http://localhost:2380 --initial-cluster-token etcd-cluster-1 --initial-cluster "etcd1=http://localhost:2380,etcd2=http://localhost:12380,etcd3=http://localhost:22380" --initial-cluster-state new
node2: etcd -name etcd2 --data-dir $PWD/node2/data --listen-client-urls http://0.0.0.0:12379 --listen-peer-urls http://0.0.0.0:12380 --advertise-client-urls http://0.0.0.0:12379 --initial-advertise-peer-urls http://localhost:12380 --initial-cluster-token etcd-cluster-1 --initial-cluster "etcd1=http://localhost:2380,etcd2=http://localhost:12380,etcd3=http://localhost:22380" --initial-cluster-state new
node3: etcd -name etcd3 --data-dir $PWD/node3/data --listen-client-urls http://0.0.0.0:22379 --listen-peer-urls http://0.0.0.0:22380 --advertise-client-urls http://0.0.0.0:22379 --initial-advertise-peer-urls http://localhost:22380 --initial-cluster-token etcd-cluster-1 --initial-cluster "etcd1=http://localhost:2380,etcd2=http://localhost:12380,etcd3=http://localhost:22380" --initial-cluster-state new
```


step4 启动
```
$ goreman start
```

step5 客户端测试
```
# V3版本(默认)
$ etcdctl set 1 1

$ etcdctl get 1
1
1

# V2版本
$ ETCDCTL_API=2 etcdctl set 1 1

$ ETCDCTL_API=2 etcdctl get 1
1
1
```