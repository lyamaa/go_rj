package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.com/go_rj/controllers"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
