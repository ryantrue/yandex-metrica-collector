package aggregator

import (
	"github.com/RyanTrue/yandex-metrica-collector/internal/collector"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAggregator_AggregateGopsutilMetrics(t *testing.T) {
	t.Run("aggregate gopsutil metrics test", func(t *testing.T) {
		c := collector.Collector
		metricsCollector := &c
		metricsAggregator := New(metricsCollector)
		metricsAggregator.AggregateGopsutilMetrics()
		assert.Equal(t, metricsCollector.GetAvailableMetrics(), []string{"FreeMemory", "TotalMemory", "CPUutilization1"})
	})
}

func TestAggregator_AggregateRuntimeMetrics(t *testing.T) {
	t.Run("aggregate runtime metrics test", func(t *testing.T) {
		c := collector.Collector
		metricsCollector := &c
		metricsAggregator := New(metricsCollector)
		metricsAggregator.AggregateRuntimeMetrics()
		assert.Equal(t, metricsCollector.GetAvailableMetrics(), []string{"Alloc", "BuckHashSys", "Frees", "GCCPUFraction", "GCSys", "HeapAlloc", "HeapIdle", "HeapInuse", "HeapObjects", "HeapReleased", "HeapSys", "Lookups", "MCacheInuse", "MCacheSys", "MSpanInuse", "MSpanSys", "Mallocs", "NextGC", "NumForcedGC", "NumGC", "OtherSys", "PauseTotalNs", "StackInuse", "StackSys", "Sys", "TotalAlloc", "RandomValue", "LastGC", "PollCount"})
	})
}
