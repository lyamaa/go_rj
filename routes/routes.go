package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.com/go_rj/controllers"
	"go.com/go_rj/middleware"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}
