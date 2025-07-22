package instance

import (
	"database/sql"
	"log/slog"
	"os"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/ncostamagna/go-posts/adapters/database"
	"github.com/ncostamagna/go-posts/internal/posts"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const METHOD = "method"

func NewPostsService(db *database.Queries, logger *slog.Logger) posts.Service {

	fieldKeys := []string{METHOD}
	repository := posts.NewRepo(db, logger)
	service := posts.NewService(logger, repository)
	return posts.NewInstrumenting(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "posts_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "posts_service",
			Name:      "request_latency_microseconds_summary",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: "api",
			Subsystem: "posts_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		service)
}

func NewDatabase() *database.Queries {
	db, err := sql.Open("postgres", os.Getenv("DB_DNS"))
	if err != nil {
		slog.Error("Error connecting to database", "error", err)
		os.Exit(1)
	}
	return database.New(db)
}
