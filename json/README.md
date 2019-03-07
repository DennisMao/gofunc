# 高性能json库 json-iterator/go

## 性能测试
先说本地和线上测试结果

win: win10 x64
```
$ go test -v --bench="."
goos: windows
goarch: amd64
pkg: auth/app/utils/json
BenchmarkDidiJsonMarshal-8               3000000               521 ns/op
BenchmarkOfficialJsonMarshal-8           3000000               599 ns/op
BenchmarkDidiJsonUnmarshal-8             3000000               421 ns/op
BenchmarkOfficialJsonUnmarshal-8         1000000              1437 ns/op
PASS
ok      auth/app/utils/json     8.071s
```

linux: amd64
```
$ go test -v --bench="."
goos: linux
goarch: amd64
pkg: testJson/json
BenchmarkDidiJsonMarshal       	 1000000	      1281 ns/op
BenchmarkOfficialJsonMarshal   	 1000000	      1623 ns/op
BenchmarkDidiJsonUnmarshal     	 1000000	      1554 ns/op
BenchmarkOfficialJsonUnmarshal 	  300000	      4664 ns/op
PASS
ok  	testJson/json	5.965s
```

序列化的差异并不算特别大,但是反序列化的性能差异很明显。

## 源码解析
