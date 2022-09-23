package routes

import (
	"github.com/abe27/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(c *fiber.App) {
	c.Get("/", controllers.HandlerHello)

	// Group Prefix Router
	r := c.Group("/api/v1")
	// User
	user := r.Group("/auth")
	user.Post("/register", controllers.Register)
	user.Post("/login", controllers.Login)
	user.Get("/me", controllers.Profile)
	user.Get("/verify", controllers.Verify)
	user.Get("/logout", controllers.Logout)
	// Area Router
	area := r.Group("/area")
	area.Get("/all", controllers.GetAllArea)
}
