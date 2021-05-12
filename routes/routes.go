package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.com/go_rj/controllers"
	"go.com/go_rj/middleware"
)

func SetUp(app *fiber.App) {

	// ROUTES FOR ALLOW ANY
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// MIDDLEWARE INITIALIZED IF USER IS AUTHENTICATED
	app.Use(middleware.IsAuthenticated)

	// ROUTS FOR USERS IF ISAUTHENTICATED
	app.Get("/api/user", controllers.User)
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/create/user", controllers.CreateUser)
	app.Get("/api/user/:id", controllers.GetUser)
	app.Put("/api/user/:id", controllers.UserUpdate)
	app.Delete("/api/user/:id", controllers.DeleteUser)
	app.Post("/api/logout", controllers.Logout)

}
