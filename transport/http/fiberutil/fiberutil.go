package fiberutil

import "github.com/gofiber/fiber/v2"

func ResponseError(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func ResponseSuccess(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"data":    data,
	})
}