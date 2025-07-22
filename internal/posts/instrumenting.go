package posts

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/google/uuid"
	"github.com/ncostamagna/go-posts/adapters/database"
)

type (
	instrumenting struct {
		requestCount          metrics.Counter
		requestLatency        metrics.Histogram
		requestLatencySummary metrics.Histogram
		s                     Service
	}

	Instrumenting interface {
		Service
	}
)

func NewInstrumenting(requestCount metrics.Counter, requestLatencySummary metrics.Histogram, requestLatency metrics.Histogram, s Service) Instrumenting {
	return &instrumenting{
		requestCount:          requestCount,
		requestLatencySummary: requestLatencySummary,
		requestLatency:        requestLatency,
		s:                     s,
	}
}

func (i *instrumenting) Store(ctx context.Context, title, content string) (*database.Post, error) {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Store").Add(1)
		i.requestLatencySummary.With("method", "Store").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "Store").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Store(ctx, title, content)
}

func (i *instrumenting) GetAll(ctx context.Context, offset, limit int32) ([]database.Post, error) {
	defer func(begin time.Time) {
		i.requestCount.With("method", "GetAll").Add(1)
		i.requestLatencySummary.With("method", "GetAll").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "GetAll").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.GetAll(ctx, offset, limit)
}

func (i *instrumenting) Get(ctx context.Context, id uuid.UUID) (*database.Post, error) {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Get").Add(1)
		i.requestLatencySummary.With("method", "Get").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "Get").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Get(ctx, id)
}

func (i *instrumenting) Delete(ctx context.Context, id uuid.UUID) error {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Delete").Add(1)
		i.requestLatencySummary.With("method", "Delete").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "Delete").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Delete(ctx, id)
}

func (i *instrumenting) Update(ctx context.Context, id uuid.UUID, title, content string) error {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Update").Add(1)
		i.requestLatencySummary.With("method", "Update").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "Update").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Update(ctx, id, title, content)
}

func (i *instrumenting) Count(ctx context.Context) (int, error) {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Count").Add(1)
		i.requestLatencySummary.With("method", "Count").Observe(time.Since(begin).Seconds())
		i.requestLatency.With("method", "Count").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.s.Count(ctx)
}
