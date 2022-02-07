package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	IsAdmin   bool
}

// SetPassword generates hashed password for user
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) CompareHashedPassword(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return err
	}
	return nil
}
