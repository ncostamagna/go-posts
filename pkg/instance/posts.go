package instance

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/ncostamagna/go-posts/adapters/database"
	"github.com/ncostamagna/go-posts/internal/posts"
	"github.com/prometheus/client_golang/prometheus"
)

func NewPostsService(db *database.Queries, logger *slog.Logger) posts.Service {
	repo := database.NewDB(db, logger)
	service := posts.NewService(logger, repo)

	requestCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "api",
		Subsystem: "posts_service",
		Name:      "request_count_total",
		Help:      "Number of requests received.",
	}, []string{"method"})

	requestLatencySummary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "api",
		Subsystem: "posts_service",
		Name:      "request_latency_seconds",
		Help:      "Total duration of requests in seconds.",
	}, []string{"method"})

	requestLatency := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "api",
		Subsystem: "posts_service",
		Name:      "request_latency_seconds",
		Help:      "Total duration of requests in seconds.",
	}, []string{"method"})

	prometheus.MustRegister(requestCount, requestLatencySummary, requestLatency)

	return posts.NewInstrumenting(requestCount, requestLatencySummary, requestLatency, service)
}

func NewDatabase() *database.Queries {
	db, err := sql.Open("postgres", os.Getenv("DB_DNS"))
	if err != nil {
		slog.Error("Error connecting to database", "error", err)
		os.Exit(1)
	}
	return database.New(db)
}
