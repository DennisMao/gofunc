# prometheus 自定义exporter编写
本demo提供了prometheus四种指标类型的exporter的设计样例。
demo参考了官方[node_exporter](https://github.com/prometheus/node_exporter)项目.


目录结构:
```
exporter/		// 工程目录,放到gopath/src下
|-- collector		// 采集器目录
|   |-- counter.go		// counter累计值类型采集器
|   |-- gauge.go		// gauge瞬时值类型采集器
|   |-- histogram.go	// histogram柱状图类型采集器
|   `-- summary.go		// summary总结类型采集器
|-- main.go			// 程序入口
`-- README.md		// 说明文档
```

exporter功能:


## exporter执行流程


## 指标类型
### counter
### gauge
### histogram
### summary

## 总结
## 参考
+ [prometheus简介](https://jeremy-xu.oschina.io/2018/08/%E7%A0%94%E7%A9%B6%E7%9B%91%E6%8E%A7%E7%B3%BB%E7%BB%9F%E4%B9%8Bprometheus/)
+ [prometheus_practice](https://github.com/songjiayang/prometheus_practice)
+ [writing_exporters](https://prometheus.io/docs/instrumenting/writing_exporters/)
