package collector

import (
	"runtime/debug"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	histogramSubsystem = "histogram"
)

// Histogram类型的一个采集接口的实现
//
// 需要分别实现 Describe 和 Collect 接口,并注册
type HistogramCollector struct {
	gc *prometheus.Desc // Gc调用的持续时间

	namespace string
}

func NewHistogramCollector(namespace string) *HistogramCollector {

	return &HistogramCollector{
		gc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, histogramSubsystem, "go_gc_duration_seconds"),
			"A summary of the GC invocation durations.",
			nil, nil),

		namespace: namespace,
	}
}

func (c *HistogramCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.gc
}

func (c *HistogramCollector) Collect(ch chan<- prometheus.Metric) {
	var stats debug.GCStats
	stats.PauseQuantiles = make([]time.Duration, 5)
	debug.ReadGCStats(&stats)

	buckets := make(map[float64]uint64)
	buckets[100] = 0
	buckets[200] = 10
	buckets[300] = 20
	buckets[600] = 20
	buckets[900] = 90

	ch <- prometheus.MustNewConstHistogram(
		c.gc,
		90,
		140,
		buckets,
	)

	return
}
