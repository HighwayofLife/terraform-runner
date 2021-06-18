package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func recordErrorMetric() {
	totalErrors.Inc()
}

var (
	totalErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tf_runner_errors_total",
		Help: "Total number of errors encountered",
	})
)
