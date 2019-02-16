package collector

import (
	"log"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/procfs"
)

const (
	counterSubsystem = "counter"
)

// Counter类型的一个采集接口的实现
//
// 需要分别实现 Describe 和 Collect 接口,并注册
type CounterCollector struct {
	cputime *prometheus.Desc // 当前程序累计允许时间

	namespace string
}

func NewCounterCollector(namespace string) *CounterCollector {

	return &CounterCollector{
		cputime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, counterSubsystem, "cputime"),
			"Process infomations' field: cputime",
			nil, nil,
		),

		namespace: namespace,
	}
}

func (c *CounterCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.cputime
}

func (c *CounterCollector) Collect(ch chan<- prometheus.Metric) {

	pid := os.Getegid()
	p, err := procfs.NewProc(pid)
	if err != nil {
		log.Println(err)
		return
	}

	s, err := p.NewStat()
	if err != nil {
		log.Println(err)
		return
	}

	cputime := s.CPUTime() // 获取到累计运行时间

	ch <- prometheus.MustNewConstMetric(
		c.cputime,
		prometheus.CounterValue,
		cputime,
	)

	return
}
