package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createCounter prometheus.Counter
	updateCounter prometheus.Counter
	deleteCounter prometheus.Counter
)

func InitMetrics() {
	createCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "roadmap_create_count_total",
		Help: "Total of create roadmap",
	})

	updateCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "roadmap_update_count_total",
		Help: "Total of update roadmap",
	})

	deleteCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "roadmap_delete_count_total",
		Help: "Total of deleted roadmap",
	})
}

func CreateCounterInc() {
	createCounter.Inc()
}

func UpdateCounterInc() {
	updateCounter.Inc()
}

func DeleteCounterInc() {
	deleteCounter.Inc()
}
