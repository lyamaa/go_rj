package main

import (
	"fmt"

	"go_rj/database"
	"go_rj/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	n1, n2 := two()
	fmt.Println(n1, n2)

	app := fiber.New()
	// CORS HEADER SET CREDENTIALS TO TRUE
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetUp(app)

	app.Listen(":8000")
}

func two() (int, int) {
	return 3, 5
}
