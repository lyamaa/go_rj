package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.com/go_rj/database"
	"go.com/go_rj/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"messgae": "Password Do Not Match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstNmae: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}
	database.DB.Create(&user)
	return c.JSON(user)
}