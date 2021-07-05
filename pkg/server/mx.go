package server

import . "github.com/prometheus/client_golang/prometheus"

type metrics struct {

}

func newMetrics(r *Registry) *metrics {
	mx := &metrics{

	}
	return mx
}
