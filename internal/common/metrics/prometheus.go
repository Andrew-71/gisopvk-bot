package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type PromMetrics struct {
	metrics map[string]prometheus.Counter
}

func NewPromMetrics() PromMetrics {
	return PromMetrics{make(map[string]prometheus.Counter)}
}

func (p PromMetrics) Inc(key string, value float64) {
	// this *feels* inefficient, but technically shouldn't be that bad
	_, ok := p.metrics[key]
	if !ok {
		p.metrics[key] = promauto.NewCounter(prometheus.CounterOpts{
			Name: key,
		})
	}
	p.metrics[key].Add(value)
}
