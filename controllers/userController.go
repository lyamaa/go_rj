package controllers

import (
	"strconv"

	"go_rj/database"
	"go_rj/models"

	"github.com/gofiber/fiber/v2"
)

// USER LIST
func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Preload("Role").Find(&users)

	return c.JSON(users)
}

// USER CREATE
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("12345678")

	database.DB.Create(&user)

	return c.JSON(user)
}

// GET USER
func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

// UPDATE USER
func UserUpdate(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)
	return c.SendStatus(204)
}
