package httpposts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewHTTPServer(endpoints Endpoints) *fiber.App {

	app := fiber.New()
	app.Use(recover.New())
	app.Get("/posts", endpoints.GetAll)
	app.Post("/posts", endpoints.Store)
	app.Get("/posts/:id", endpoints.Get)
	app.Patch("/posts/:id", endpoints.Update)
	app.Delete("/posts/:id", endpoints.Delete)
	app.Get("/metrics", promhttp.Handler())
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	return app
}