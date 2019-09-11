# goexperience
goexperience 是自己平时使用的Go库使用demo和经验,还有部分自己学习过程中的设计项目。  
使用说明:  
>从Go 1.11开始,本工程下的涉及import的项目都采用`go mod`来管理,各位需要学习的同学可以先查看项目下的go mod文件查看版本。对于官方库的DEMO 默认按照最新版本(Go官方库本身有兼容性),如发现与新版有冲突无法编译的,请提[issue](https://github.com/DennisMao/goexperience/issues/new)告知,本人会尽快更新。希望本项目的DEMO能帮助大家,愿与各位Gopher共同学习,一起成长。


## 业务相关

## 官方库
|库名称|功能|最新版本|Demo|
|:-|:-|:-|:-|
|syscall|调用动态链接库|-|[Demo](/loaddll)|
|plugin|调用动态链接库|-|[Demo](/plugin)|
|testing|单元测试|-|[Demo](/unittest)|
|testing|性能压测|-|[Demo](/benchmark)|
|goroutine|协程并发控制|-|[Demo](/concurrency)
|compress/gzip|Gzip压缩解压|-|[Demo](/compress/gzip)|
|sync/cond|协程阻塞通知|-|[Demo](/sync/cond)|
|go/scanner|Go代码词法分析库|-|[Demo](/gosrc/scanner)|
|go/parser|Go代码语法分析库|-|[Demo](/gosrc/parser)|
|go/printer|Go代码生成库|-|[Demo](/gosrc/printer)|
|net/http|http网络库|-|[Demo](/net/http)|
|sync/map|并发安全map|-|[Demo](/sync/map)|
|sync/atomic|原子操作|-|[Demo](/sync/atomic)|
|runtime/trace|trace库|-|[Demo](/runtime/trace)|
|runtime/pprof|pprof库|-|[Demo](/runtime/pprof)|

## 开源库

|库名称|功能|Demo版本|Demo|
|:-|:-|:-|:-|
|[xlsx](https://github.com/tealeg/xlsx)|Excel解析库|-|[Demo]()|
|[gRPC](https://github.com/grpc/go-grpc)|支持多语言的RPC库|-|[Demo](/rpc)|
|[plot](https://gonum.org/v1/plot)|类似mathlab的polt曲线图生成工具|-|[Demo](/plot)|
|[etcd/raft](https://github.com/etcd-io/etcd/raft)|分布式一致性共识算法实现的简易kv数据库|-|[Demo](/raft/raft-example)|
|[google/btree](https://github.com/google/btree)|B树的Go实现|-|[Demo](/google/btree-example)|
|[elastic.v5](https://github.com/olivere/elastic)|elasticsearch链接库|-|[Demo](/elasticsearch/README.md)|
|[jaeger/client]()|jaeger的opentracing使用|-|[Demo](/jaeger/testUdpSender)|
|[prometheus/exporter]()|prometheus的exporter编写|-|[Demo](/prometheus/exporter/README.md)|
|[golang-lru]()|cache lru算法包使用|-|[Demo](/cache/README.md)|
|[etcd-client]()|etcd客户端 v2 v3版本|-|[Demo](/etcd/README.md)|
|[json-iterator/go](github.com/json-iterator/go)|高性能json库|-|[Demo](/json/README.md)|
|[radix]()|redis集群客户端|-|[Demo](/redis/radix)|
|[redigo/redis]()|redis单点&哨兵客户端|-|[Demo](/redis/redigo)|
|[kafka/sarama]()|kafka集群客户端|-|[Demo](/kafka/sarama)|
|[mysql/canal]()|mysql Binlog协议库|-|[Demo](/mysql_dm/canal)|
|[mysql/dump]()|mysql 热迁移dump协议库|-|[Demo](/mysql_dm/dump)|
|[mysql/sqlparser]|sql解析库(以mysql为主)|-|[Demo](/mysql_dm/sqlparser)|


## 学习

|项目名称|功能|Demo版本|Demo|
|:-|:-|:-|:-|
|[redisgo]()|Redis的go实现|-|[Demo](/redisgo)|
|[da/datastructures]()|数据结构的go实现|-|[Code](/data/datastructures)|
|[da/alogrithms]()|通用算法的go实现|-|[Code](/data/alogrithms)|

### DL/ML/CV

### 数据结构
|项目名称|功能|Demo版本|Demo|
|:-|:-|:-|:-|
|[da/datastructures]()|数据结构的go实现|-|[Code](/data/datastructures)|

### 算法
|项目名称|功能|Demo版本|Demo|
|:-|:-|:-|:-|
|[da/alogrithms]()|通用算法的go实现|-|[Code](/data/alogrithms)|

