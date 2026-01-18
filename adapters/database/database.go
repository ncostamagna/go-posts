package database

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type (
	Database interface {
		GetAll(ctx context.Context, offset, limit int32) ([]Post, error)
		Get(ctx context.Context, id uuid.UUID) (*Post, error)
		Delete(ctx context.Context, id uuid.UUID) error
		Update(ctx context.Context, id uuid.UUID, title, content string) error
		Count(ctx context.Context) (int, error)
		Store(ctx context.Context, post *Post) error
	}

	db struct {
		q  *Queries
		log *slog.Logger
	}
)

func NewDB(q *Queries, l *slog.Logger) Database {
	return &db{
		q:  q,
		log: l,
	}
}

func (d *db) Store(ctx context.Context, post *Post) error {
	id, err := d.q.InsertPost(ctx, InsertPostParams{
		Title:   post.Title,
		Content: post.Content,
	})
	if err != nil {
		return err
	}
	post.ID = id
	return nil
}

func (d *db) GetAll(ctx context.Context, offset, limit int32) ([]Post, error) {
	return d.q.GetAllPosts(ctx, GetAllPostsParams{
		Offset: offset,
		Limit:  limit,
	})
}

func (d *db) Get(ctx context.Context, id uuid.UUID) (*Post, error) {
	post, err := d.q.GetPostById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (d *db) Delete(ctx context.Context, id uuid.UUID) error {
	return d.q.DeletePost(ctx, id)
}

func (d *db) Update(ctx context.Context, id uuid.UUID, title, content string) error {
	return d.q.UpdatePost(ctx, UpdatePostParams{
		ID:      id,
		Title:   title,
		Content: content,
	})
}

func (d *db) Count(ctx context.Context) (int, error) {
	count, err := d.q.CountPosts(ctx)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
