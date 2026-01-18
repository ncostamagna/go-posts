package main

import (
	"time"

	"github.com/ncostamagna/go-posts/pkg/instance"

	"flag"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ncostamagna/go-posts/pkg/log"
	"github.com/ncostamagna/go-posts/transport/http/httpposts"
)

const writeTimeout = 10 * time.Second
const readTimeout = 4 * time.Second
const defaultURL = "0.0.0.0:80"

func main() {

	_ = godotenv.Load()

	logger := log.New(log.Config{
		AppName:   "posts-service",
		Level:     os.Getenv("LOG_LEVEL"),
		AddSource: true,
	})

	defer func() {
		if r := recover(); r != nil {
			logger.Error("Application panicked", "err", r)
		}
	}()

	flag.Parse()

	postsSrv := instance.NewPostsService(instance.NewDatabase(), logger)

	pagLimDef := "30"

	app := httpposts.NewHTTPServer(httpposts.MakePostsEndpoints(postsSrv, httpposts.Config{LimPageDef: pagLimDef}))

	url := os.Getenv("APP_URL")
	if url == "" {
		url = defaultURL
	}

	errs := make(chan error)

	go func() {
		logger.Info("Listening", "url", url)
		errs <- app.Listen(url)
	}()

	err := <-errs
	if err != nil {
		logger.Error("Error server", "err", err)
	}

}
