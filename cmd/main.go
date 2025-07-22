package main

import (
	"time"

	"github.com/ncostamagna/go-posts/pkg/instance"

	"context"
	"flag"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ncostamagna/go-posts/internal/posts"
	"github.com/ncostamagna/go-posts/pkg/log"
	"github.com/ncostamagna/go-posts/transport/httpTransport"
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
	ctx := context.Background()

	postsSrv := instance.NewPostsService(instance.NewDatabase(), logger)

	pagLimDef := "30"

	h := httpTransport.NewHTTPServer(ctx, posts.MakeEndpoints(postsSrv, posts.Config{LimPageDef: pagLimDef}))

	url := os.Getenv("APP_URL")
	if url == "" {
		url = defaultURL
	}

	srv := &http.Server{
		Handler:      accessControl(h),
		Addr:         url,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	errs := make(chan error)

	go func() {
		logger.Info("Listening", "url", url)
		errs <- srv.ListenAndServe()
	}()

	err := <-errs
	if err != nil {
		logger.Error("Error server", "err", err)
	}

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With")

		if r.Method == http.MethodOptions {
			return
		}

		h.ServeHTTP(w, r)
	})
}
