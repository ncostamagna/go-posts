package posts

import (
	"context"
	"log/slog"

	"github.com/ncostamagna/go-posts/adapters/database"
)

type (
	Filters struct {
		Name string
	}

	Service interface {
		Store(ctx context.Context, title, content string) (*database.Post, error)
	}

	service struct {
		log  *slog.Logger
		repo Repository
	}
)

func NewService(l *slog.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) Store(ctx context.Context, title, content string) (*database.Post, error) {

	post := &database.Post{
		Title:   title,
		Content: content,
	}

	if err := s.repo.Store(ctx, post); err != nil {
		return nil, err
	}
	s.log.Info("post stored", "post id", post.ID)
	return post, nil
}