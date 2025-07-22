package posts

import (
	"context"

	"github.com/ncostamagna/go-http-utils/response"
	intPosts "github.com/ncostamagna/go-posts/internal/posts"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type (
	Endpoints struct {
		Get    endpoint.Endpoint
		GetAll endpoint.Endpoint
		Store  endpoint.Endpoint
		Update endpoint.Endpoint
		Delete endpoint.Endpoint
	}

	StoreReq struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	GetReq struct {
		ID uuid.UUID `json:"id"`
	}

	GetAllReq struct {
		Limit int32
		Page  int32
	}

	UpdateReq struct {
		ID      uuid.UUID `json:"id"`
		Title   string    `json:"title"`
		Content string    `json:"content"`
	}

	DeleteReq struct {
		ID uuid.UUID `json:"id"`
	}

	Config struct {
		LimPageDef string
	}
)

func MakeEndpoints(s intPosts.Service, c Config) Endpoints {
	return Endpoints{
		Get:    makeGet(s),
		GetAll: makeGetAll(s, c),
		Store:  makeStore(s),
		Update: makeUpdate(s),
		Delete: makeDelete(s),
	}
}

func makeGet(service intPosts.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetReq)

		post, err := service.Get(ctx, req.ID)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("Success", post, nil), nil
	}
}

func makeGetAll(service intPosts.Service, _ Config) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllReq)

		posts, err := service.GetAll(ctx, req.Page, req.Limit)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("Success", posts, nil), nil
	}
}

func makeStore(service intPosts.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StoreReq)

		if req.Title == "" {
			return nil, response.BadRequest("Title is required")
		}

		if req.Content == "" {
			return nil, response.BadRequest("Content is required")
		}

		post, err := service.Store(ctx, req.Title, req.Content)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}
		return response.Created("Success", post, nil), nil
	}
}

func makeUpdate(service intPosts.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateReq)

		if req.ID == uuid.Nil {
			return nil, response.BadRequest("ID is required")
		}

		if req.Title == "" {
			return nil, response.BadRequest("Title is required")
		}

		if req.Content == "" {
			return nil, response.BadRequest("Content is required")
		}

		if err := service.Update(ctx, req.ID, req.Title, req.Content); err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("Success", nil, nil), nil
	}
}

func makeDelete(service intPosts.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteReq)

		if req.ID == uuid.Nil {
			return nil, response.BadRequest("ID is required")
		}

		if err := service.Delete(ctx, req.ID); err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("Success", nil, nil), nil
	}
}
