package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(router fiber.Router) {
	user := router.Group("/users")
	user.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("User route")
	})
}
