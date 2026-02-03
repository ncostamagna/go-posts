package posts

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ncostamagna/go-posts/adapters/database"
	"github.com/prometheus/client_golang/prometheus"
)

type (
	instrumenting struct {
		requestCount          *prometheus.CounterVec
		requestLatency        *prometheus.HistogramVec
		requestLatencySummary *prometheus.SummaryVec
		s                     Service
	}

	Instrumenting interface {
		Service
	}
)

func NewInstrumenting(requestCount *prometheus.CounterVec, requestLatencySummary *prometheus.SummaryVec, requestLatency *prometheus.HistogramVec, s Service) Instrumenting {
	return &instrumenting{
		requestCount:          requestCount,
		requestLatencySummary: requestLatencySummary,
		requestLatency:        requestLatency,
		s:                     s,
	}
}

func (i *instrumenting) Store(ctx context.Context, title, content string) (*database.Post, error) {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("Store").Inc()
		i.requestLatencySummary.WithLabelValues("Store").Observe(time.Since(begin).Seconds())
		i.requestLatency.WithLabelValues("Store").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Store(ctx, title, content)
}

func (i *instrumenting) GetAll(ctx context.Context, offset, limit int32) ([]database.Post, error) {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("GetAll").Inc()
		i.requestLatencySummary.WithLabelValues("GetAll").Observe(time.Since(begin).Seconds())
		i.requestLatency.WithLabelValues("GetAll").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.GetAll(ctx, offset, limit)
}

func (i *instrumenting) Get(ctx context.Context, id uuid.UUID) (*database.Post, error) {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("Get").Inc()
		i.requestLatencySummary.WithLabelValues("Get").Observe(time.Since(begin).Seconds())
		i.requestLatency.WithLabelValues("Get").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Get(ctx, id)
}

func (i *instrumenting) Delete(ctx context.Context, id uuid.UUID) error {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("Delete").Inc()
		i.requestLatencySummary.WithLabelValues("Delete").Observe(time.Since(begin).Seconds())
		i.requestLatency.WithLabelValues("Delete").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Delete(ctx, id)
}

func (i *instrumenting) Update(ctx context.Context, id uuid.UUID, title, content string) error {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("Update").Inc()
		i.requestLatencySummary.WithLabelValues("Update").Observe(time.Since(begin).Seconds())
		i.requestLatency.WithLabelValues("Update").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Update(ctx, id, title, content)
}

func (i *instrumenting) Count(ctx context.Context) (int, error) {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("Count").Inc()
		i.requestLatencySummary.WithLabelValues("Count").Observe(time.Since(begin).Seconds())
		i.requestLatency.WithLabelValues("Count").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Count(ctx)
}
