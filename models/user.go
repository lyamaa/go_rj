package models

import "golang.org/x/crypto/bcrypt"

// USER STRUCTURE
type User struct {
	Id        uint   `json:"id"`
	FirstNmae string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}

// Password HASHING
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user.Password = hashedPassword
}

// COMPARE PASSWORD
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
