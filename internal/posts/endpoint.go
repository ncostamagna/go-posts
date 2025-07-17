package posts

import (
	"context"

	"github.com/ncostamagna/go-http-utils/response"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	// Endpoints struct.
	Endpoints struct {
		Get    Controller
		GetAll Controller
		Store  Controller
		Update Controller
		Delete Controller
	}

	StoreReq struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
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

func MakeEndpoints(s Service, c Config) Endpoints {
	return Endpoints{
		Get:    makeGet(s),
		GetAll: makeGetAll(s, c),
		Store:  makeStore(s),
		Update: makeUpdate(s),
		Delete: makeDelete(s),
	}
}

func makeGet(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		
		return response.OK("Success", nil,  nil), nil
	}
}

func makeGetAll(service Service, c Config) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		return response.OK("Success", nil, nil), nil
	}
}

func makeStore(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		return response.Created("Success", nil, nil), nil
	}
}

func makeUpdate(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		return response.OK("Success", "UPDATE: testing 1234 6789", nil), nil
	}
}

func makeDelete(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		return response.OK("Success", "", nil), nil
	}
}
