package httpposts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewHTTPServer(endpoints Endpoints) *fiber.App {

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{                                                              
		AllowOrigins: "*",                                                                     
		AllowMethods: "GET, POST, OPTIONS, HEAD",                                              
		AllowHeaders: "Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With",                                              
	})) 
	app.Get("/posts", endpoints.GetAll)
	app.Post("/posts", endpoints.Store)
	app.Get("/posts/:id", endpoints.Get)
	app.Patch("/posts/:id", endpoints.Update)
	app.Delete("/posts/:id", endpoints.Delete)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	return app
}