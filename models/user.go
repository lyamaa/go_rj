package models

type User struct {
	Id        uint
	FirstNmae string
	LastName  string
	Email     string `gorm:"unique"`
	Password  []byte
}
