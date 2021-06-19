package routes

import (
	"go_rj/controllers"
	"go_rj/middleware"

	"github.com/gofiber/fiber/v2"
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

	// ROLE ROUTES
	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/role", controllers.CreateRole)
	app.Get("/api/role/:id", controllers.GetRole)
	app.Put("/api/role/:id", controllers.UpdateRole)
	app.Delete("/api/role/:id", controllers.DeleteRole)

	// Permissions Routes
	app.Get("/api/permissions", controllers.AllPermissions)

	app.Post("/api/logout", controllers.Logout)

}
