package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.com/go_rj/database"
	"go.com/go_rj/routes"
)

func main() {
	database.Connect()
	n1, n2 := two()
	fmt.Println(n1, n2)

	app := fiber.New()

	routes.SetUp(app)

	app.Listen(":3000")
}

func two() (int, int) {
	return 3, 5
}
