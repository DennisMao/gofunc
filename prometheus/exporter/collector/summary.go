package collector

import (
	"runtime/debug"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	summarySubsystem = "summary"
)

// Summary类型的一个采集接口的实现
//
// 需要分别实现 Describe 和 Collect 接口,并注册
type SummaryCollector struct {
	gc *prometheus.Desc // Gc调用的持续时间

	namespace string
}

func NewSummaryCollector(namespace string) *SummaryCollector {

	return &SummaryCollector{
		gc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, summarySubsystem, "gc_duration_seconds"),
			"A summary of the GC invocation durations.",
			nil, nil),

		namespace: namespace,
	}
}

func (c *SummaryCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.gc
}

func (c *SummaryCollector) Collect(ch chan<- prometheus.Metric) {

	var stats debug.GCStats
	stats.PauseQuantiles = make([]time.Duration, 5)
	debug.ReadGCStats(&stats)

	quantiles := make(map[float64]float64)
	for idx, pq := range stats.PauseQuantiles[1:] {
		quantiles[float64(idx+1)/float64(len(stats.PauseQuantiles)-1)] = pq.Seconds() // 分值计算
	}
	quantiles[0.0] = stats.PauseQuantiles[0].Seconds()
	ch <- prometheus.MustNewConstSummary(
		c.gc,
		uint64(stats.NumGC),
		stats.PauseTotal.Seconds(),
		quantiles,
	)

	return
}
