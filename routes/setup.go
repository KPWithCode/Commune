package router

import (
    "github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
    return c.SendString("Hello World!")
}

func helloBlog(c *fiber.Ctx) error {
    return c.SendString("Hello Blog!!!!")
}

//==============SetupRoutes================
func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")
	api.Get("/blog",helloBlog)
    api.Get("/", hello)
}