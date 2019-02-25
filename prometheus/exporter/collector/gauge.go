package collector

import (
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	gaugeSubsystem = "gauge"
)

// Gauge类型的一个采集接口的实现
//
// 需要分别实现 Describe 和 Collect 接口,并注册
type GaugeCollector struct {
	goroutines *prometheus.Desc

	namespace string
}

func NewGaugeCollector(namespace string) *GaugeCollector {

	return &GaugeCollector{
		goroutines: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, gaugeSubsystem, "goroutines"),
			"Go runtime infomations' field: goroutines",
			nil, nil,
		),

		namespace: namespace,
	}
}

func (c *GaugeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.goroutines
}

func (c *GaugeCollector) Collect(ch chan<- prometheus.Metric) {

	goroutinesValue := runtime.NumGoroutine()

	ch <- prometheus.MustNewConstMetric(
		c.goroutines,
		prometheus.GaugeValue,
		float64(goroutinesValue),
	)

	return
}
