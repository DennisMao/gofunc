package main

import (
	"flag"
	"log"
	"net/http"

	collector "nginx-log-exporter/collector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr, namespace string
)

func main() {
	flag.StringVar(&addr, "web.listen-address", ":9999", "Address to listen on for the web interface and API.")
	flag.StringVar(&namespace, "namespace", "app", "Namespace for currenct exporter.")

	flag.Parse()

	// 新建注册器
	// 通过新建的注册器进行采集器绑定，将不再默认绑定"go底层运行指标"和"进程指标"信息
	// 但通过http请求"/metrics"信息的指标还是会有。
	// 此处可参考代码 https://github.com/prometheus/client_golang/blob/4c99dd66303a54cbf8559dd6110d5c30b1819e4c/prometheus/registry.go#L59
	//
	// 如需保留默认绑定的指标信息,可替换为以下代码:
	//
	//   prometheus.MustRegister(collector.NewGaugeCollector(namespace))     // Gauge 瞬时值类型采集器
	//   prometheus.MustRegister(collector.NewCounterCollector(namespace))   // Counter 累计值类型采集器
	//   prometheus.MustRegister(collector.NewSummaryCollector(namespace))   // Summary 采样结果采集器
	//   prometheus.MustRegister(collector.NewHistogramCollector(namespace)) // Histogram 采样值采集器
	//
	//   http.Handle("/metrics",prometheus.Handler())
	//
	register := prometheus.NewRegistry()
	// 注册采集器
	register.MustRegister(collector.NewGaugeCollector(namespace))     // Gauge 瞬时值类型采集器
	register.MustRegister(collector.NewCounterCollector(namespace))   // Counter 累计值类型采集器
	register.MustRegister(collector.NewSummaryCollector(namespace))   // Summary 采样结果采集器
	register.MustRegister(collector.NewHistogramCollector(namespace)) // Histogram 采样值采集器

	prometheus.DefaultRegisterer = register

	log.Printf("Running exporter node server on address %s\n", addr)
	http.Handle("/metrics", promhttp.InstrumentMetricHandler(register, promhttp.HandlerFor(prometheus.Gatherer(register), promhttp.HandlerOpts{})))

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
