package api

import "github.com/gofiber/fiber/v2"

func Hello(r fiber.Router) {
	r.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})
}
