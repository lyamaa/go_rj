package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.com/go_rj/database"
	"go.com/go_rj/models"
	"go.com/go_rj/utils"
)

// USER REGISTER
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

	user := models.User{
		FirstNmae: data["first_name"],
		LastName:  data["last_name"],
		Username:  data["username"],
		Email:     data["email"],
	}
	user.SetPassword(data["password"])
	database.DB.Create(&user)
	return c.JSON(user)
}

// USER LOGIN
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email =?", data["email"]).First(&user)
	// Check User is valid
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "No User Found...",
		})
	}
	// Compare password
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password does not match...",
		})
	}
	// strconv.Itoa(int(user.Id)),

	// JWT initializations
	token, err := utils.GenerateJWT(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
