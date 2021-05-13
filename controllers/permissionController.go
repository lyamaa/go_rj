package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.com/go_rj/database"
	"go.com/go_rj/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
