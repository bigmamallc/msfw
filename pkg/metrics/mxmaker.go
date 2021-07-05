package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	r           *prometheus.Registry
	subsystem   string
	ConstLabels prometheus.Labels
}

func NewMetrics(r *prometheus.Registry, subsystem string) *Metrics {
	return &Metrics{
		r:         r,
		subsystem: subsystem,
	}
}

func (m *Metrics) Counter(name string) prometheus.Counter {
	c := prometheus.NewCounter(prometheus.CounterOpts{
		Subsystem:   m.subsystem,
		Name:        name,
		ConstLabels: m.ConstLabels,
	})
	m.r.MustRegister(c)
	return c
}

func (m *Metrics) CounterVec(name string, labels []string) *prometheus.CounterVec {
	cv := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: m.subsystem,
		Name:      name,
		ConstLabels: m.ConstLabels,
	}, labels)
	m.r.MustRegister(cv)
	return cv
}

func (m *Metrics) Histogram(name string) prometheus.Histogram {
	h := prometheus.NewHistogram(prometheus.HistogramOpts{
		Subsystem:   m.subsystem,
		Name:        name,
		ConstLabels: m.ConstLabels,
	})
	m.r.MustRegister(h)
	return h
}

func (m *Metrics) HistogramWithBuckets(name string, buckets []float64) prometheus.Histogram {
	h := prometheus.NewHistogram(prometheus.HistogramOpts{
		Subsystem:   m.subsystem,
		Name:        name,
		Buckets:     buckets,
		ConstLabels: m.ConstLabels,
	})
	m.r.MustRegister(h)
	return h
}

func (m *Metrics) HistogramVecWithBuckets(name string, labels []string, buckets []float64) *prometheus.HistogramVec {
	hv := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: m.subsystem,
		Name:      name,
		Buckets:   buckets,
		ConstLabels: m.ConstLabels,
	}, labels)
	m.r.MustRegister(hv)
	return hv
}

func (m *Metrics) HistogramVec(name string, labels []string) *prometheus.HistogramVec {
	return m.HistogramVecWithBuckets(name, labels, nil)
}


func (m *Metrics) GaugeVec(name string, labels[] string) *prometheus.GaugeVec {
	gv := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: m.subsystem,
		Name:      name,
		ConstLabels: m.ConstLabels,
	}, labels)
	m.r.MustRegister(gv)
	return gv
}

func (m *Metrics) Gauge(name string) prometheus.Gauge {
	g := prometheus.NewGauge(prometheus.GaugeOpts{
		Subsystem: m.subsystem,
		Name:      name,
		ConstLabels: m.ConstLabels,
	})
	m.r.MustRegister(g)
	return g
}
