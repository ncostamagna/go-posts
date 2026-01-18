package httpposts

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ncostamagna/go-posts/internal/posts"
	"fmt"
	"github.com/google/uuid"
	"github.com/ncostamagna/go-posts/transport/http/fiberutil"
)

type (
	Endpoints struct {
		Get    fiber.Handler
		GetAll fiber.Handler
		Store  fiber.Handler
		Update fiber.Handler
		Delete fiber.Handler
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

func MakePostsEndpoints(s posts.Service, c Config) Endpoints {
	return Endpoints{
		Get:    makeGet(s),
		GetAll: makeGetAll(s, c),
		Store:  makeStore(s),
		Update: makeUpdate(s),
		Delete: makeDelete(s),
	}
}

func makeGet(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := GetReq{
			ID: uuid.MustParse(c.Params("id")),
		}

		post, err := service.Get(c.Context(), req.ID)
		if err != nil {
			return fiberutil.ResponseError(c, fiber.StatusInternalServerError, err)
		}

		return fiberutil.ResponseSuccess(c, fiber.StatusOK, post)
	}
}

func makeGetAll(service posts.Service, _ Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := GetAllReq{
			Page:  int32(c.QueryInt("page")),
			Limit: int32(c.QueryInt("limit")),
		}
		fmt.Println(req)

		posts, err := service.GetAll(c.Context(), req.Page, req.Limit)
		if err != nil {
			return fiberutil.ResponseError(c, fiber.StatusInternalServerError, err)
		}

		return fiberutil.ResponseSuccess(c, fiber.StatusOK, posts)
	}
}

func makeStore(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := StoreReq{}
		if err := c.BodyParser(&req); err != nil {
			return fiberutil.ResponseError(c, fiber.StatusBadRequest, err)
		}

		if req.Title == "" {
			return fiberutil.ResponseError(c, fiber.StatusBadRequest, errors.New("title is required"))
		}

		if req.Content == "" {
			return fiberutil.ResponseError(c, fiber.StatusBadRequest, errors.New("content is required"))
		}

		post, err := service.Store(c.Context(), req.Title, req.Content)
		if err != nil {
			return fiberutil.ResponseError(c, fiber.StatusInternalServerError, err)
		}
		return fiberutil.ResponseSuccess(c, fiber.StatusCreated, post)
	}
}

func makeUpdate(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := UpdateReq{}
		if err := c.BodyParser(&req); err != nil {
			return fiberutil.ResponseError(c, fiber.StatusBadRequest, err)
		}

		if req.ID == uuid.Nil {
			return fiberutil.ResponseError(c, fiber.StatusBadRequest, errors.New("id is required"))
		}

		if req.Title == "" {
			return fiberutil.ResponseError(c, fiber.StatusBadRequest, errors.New("title is required"))
		}

		if req.Content == "" {
			return fiberutil.ResponseError(c, fiber.StatusBadRequest, errors.New("content is required"))
		}

		if err := service.Update(c.Context(), req.ID, req.Title, req.Content); err != nil {
			return fiberutil.ResponseError(c, fiber.StatusInternalServerError, err)
		}

		return fiberutil.ResponseSuccess(c, fiber.StatusOK, nil)
	}
}

func makeDelete(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := DeleteReq{
			ID: uuid.MustParse(c.Params("id")),
		}

		if req.ID == uuid.Nil {
			return fiberutil.ResponseError(c, fiber.StatusBadRequest, errors.New("id is required"))
		}

		if err := service.Delete(c.Context(), req.ID); err != nil {
			return fiberutil.ResponseError(c, fiber.StatusInternalServerError, err)
		}

		return fiberutil.ResponseSuccess(c, fiber.StatusOK, nil)
	}
}
