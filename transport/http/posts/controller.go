package posts

import (
	"context"

	intPosts "github.com/ncostamagna/go-posts/internal/posts"
	"github.com/ncostamagna/go-http-utils/response"

	"github.com/go-kit/kit/endpoint"
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
		ID int `json:"productId"`
	}

	GetAllReq struct {
		Name  string
		Limit int
		Page  int
	}

	UpdateReq struct {
		ID          int
		Name        *string  `json:"name"`
		Description *string  `json:"description"`
		Price       *float64 `json:"price"`
	}

	DeleteReq struct {
		ID int
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
		
		return response.OK("Success", nil,  nil), nil
	}
}

func makeGetAll(service intPosts.Service, c Config) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		return response.OK("Success", nil, nil), nil
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

		return response.OK("Success", "UPDATE: testing 1234 6789", nil), nil
	}
}

func makeDelete(service intPosts.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		return response.OK("Success", "", nil), nil
	}
}
