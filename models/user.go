package models

type User struct {
	Id        uint
	FirstNmae string
	LastName  string
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  []byte
}
