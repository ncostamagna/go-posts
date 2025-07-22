package posts

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/ncostamagna/go-posts/adapters/database"
)

type (
	Repository interface {
		GetAll(ctx context.Context, offset, limit int32) ([]database.Post, error)
		Get(ctx context.Context, id uuid.UUID) (*database.Post, error)
		Delete(ctx context.Context, id uuid.UUID) error
		Update(ctx context.Context, id uuid.UUID, title, content string) error
		Count(ctx context.Context) (int, error)
		Store(ctx context.Context, post *database.Post) error
	}

	repo struct {
		db  *database.Queries
		log *slog.Logger
	}
)

func NewRepo(db *database.Queries, l *slog.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Store(ctx context.Context, post *database.Post) error {
	id, err := r.db.InsertPost(ctx, database.InsertPostParams{
		Title:   post.Title,
		Content: post.Content,
	})
	if err != nil {
		return err
	}
	post.ID = id
	return nil
}

func (r *repo) GetAll(ctx context.Context, offset, limit int32) ([]database.Post, error) {
	return r.db.GetAllPosts(ctx, database.GetAllPostsParams{
		Offset: offset,
		Limit:  limit,
	})
}

func (r *repo) Get(ctx context.Context, id uuid.UUID) (*database.Post, error) {
	post, err := r.db.GetPostById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *repo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.DeletePost(ctx, id)
}

func (r *repo) Update(ctx context.Context, id uuid.UUID, title, content string) error {
	return r.db.UpdatePost(ctx, database.UpdatePostParams{
		ID:      id,
		Title:   title,
		Content: content,
	})
}

func (r *repo) Count(ctx context.Context) (int, error) {
	count, err := r.db.CountPosts(ctx)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
