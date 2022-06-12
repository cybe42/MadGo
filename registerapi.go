package main

import (
	"github.com/cybe42/MadGo/api"
	"github.com/gofiber/fiber/v2"
)

func loadAPI(r fiber.Router) { // Call all your API's in this function
	api.Hello(r)
}

func groupAPIHandler(c *fiber.Ctx) error {
	return c.Next() // You can modify this API Group handler middleware.
}

func StartAPI(app *fiber.App) {
	APIGroup := app.Group("/api", groupAPIHandler)
	loadAPI(APIGroup)
}
