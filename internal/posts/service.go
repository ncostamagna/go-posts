package posts

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/ncostamagna/go-posts/adapters/database"
)

type (
	Filters struct {
		Name string
	}

	Service interface {
		Store(ctx context.Context, title, content string) (*database.Post, error)
		GetAll(ctx context.Context, offset, limit int32) ([]database.Post, error)
		Get(ctx context.Context, id uuid.UUID) (*database.Post, error)
		Delete(ctx context.Context, id uuid.UUID) error
		Update(ctx context.Context, id uuid.UUID, title, content string) error
		Count(ctx context.Context) (int, error)
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
		s.log.Error("error storing post", "error", err)
		return nil, err
	}
	s.log.Info("post stored", "post id", post.ID)
	return post, nil
}

func (s service) GetAll(ctx context.Context, offset, limit int32) ([]database.Post, error) {
	posts, err := s.repo.GetAll(ctx, offset, limit)
	if err != nil {
		s.log.Error("error fetching posts", "error", err)
		return nil, err
	}
	s.log.Info("posts fetched", "posts", len(posts))
	return posts, nil
}

func (s service) Get(ctx context.Context, id uuid.UUID) (*database.Post, error) {
	post, err := s.repo.Get(ctx, id)
	if err != nil {
		s.log.Error("error fetching post", "error", err)
		return nil, err
	}
	s.log.Info("post fetched", "post id", post.ID)
	return post, nil
}

func (s service) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		s.log.Error("error deleting post", "error", err)
		return err
	}
	s.log.Info("post deleted", "post id", id)
	return nil
}

func (s service) Update(ctx context.Context, id uuid.UUID, title, content string) error {
	err := s.repo.Update(ctx, id, title, content)
	if err != nil {
		s.log.Error("error updating post", "error", err)
		return err
	}
	s.log.Info("post updated", "post id", id)
	return nil
}

func (s service) Count(ctx context.Context) (int, error) {
	count, err := s.repo.Count(ctx)
	if err != nil {
		s.log.Error("error counting posts", "error", err)
		return 0, err
	}
	s.log.Info("posts counted", "count", count)
	return count, nil
}
