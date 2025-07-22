package posts

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
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