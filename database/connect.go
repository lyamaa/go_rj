package database

import (
	"fmt"

	"go_rj/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DATABSE CONNECTION METHOD
func Connect() {
	database, err := gorm.Open(mysql.Open("root:root@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	} else {
		fmt.Println("Connecting..... wait!!")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})

}
