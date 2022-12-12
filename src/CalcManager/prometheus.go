package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// we create a new custom metric of type counter
var userStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_user_status_count", // metric name
		Help: "Count of status returned by user.",
	},
	[]string{"user", "status"}, // labels
)

func init() {
	// we need to register the counter so prometheus can collect this metric
	prometheus.MustRegister(userStatus)
}

type MyRequest struct {
	User string
}

