package posts

import "log/slog"

type (
	Repository interface {
	}

	repo struct {
		log *slog.Logger
	}
)

/*
import (
	"context"
	"log/slog"
	"maps"
	"slices"

	"github.com/ncostamagna/go-posts/internal/domain"
)

type (
	Repository interface {
		Store(ctx context.Context, product *domain.Post) error
		GetAll(ctx context.Context, offset, limit int) ([]domain.Post, error)
		Get(ctx context.Context, id int) (*domain.Post, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, id int, title, content *string) error
		Count(ctx context.Context) (int, error)
	}

	db struct {
		products map[int]domain.Product
		maxID    int
	}
	repo struct {
		db  db
		log *slog.Logger
	}
)

// NewRepo is a repositories handler.
func NewRepo(l *slog.Logger) Repository {
	return &repo{
		db: db{
			products: make(map[int]domain.Product),
			maxID:    0,
		},
		log: l,
	}
}

func (r *repo) Store(_ context.Context, product *domain.Product) error {
	r.db.maxID++
	product.ID = r.db.maxID
	r.db.products[r.db.maxID] = *product
	return nil
}

func (r *repo) GetAll(_ context.Context, _, _ int) ([]domain.Product, error) {
	return slices.Collect(maps.Values(r.db.products)), nil
}

func (r *repo) Get(_ context.Context, id int) (*domain.Product, error) {

	prod, ok := r.db.products[id]
	if !ok {
		r.log.Error("product not found", "id", id)
		return nil, ErrNotFound{id}
	}

	return &prod, nil
}

func (r *repo) Delete(_ context.Context, id int) error {
	delete(r.db.products, id)
	return nil
}

func (r *repo) Update(ctx context.Context, id int, name, description *string, price *float64) error {
	p, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	if name != nil {
		p.Name = *name
	}

	if description != nil {
		p.Description = *description
	}

	if price != nil {
		p.Price = *price
	}

	return nil
}

func (r *repo) Count(_ context.Context) (int, error) {
	return r.db.maxID, nil
}
*/