package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	BaseModel
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" gorm:"unique"`
	Password    []byte `json:"-"`
	IsAmbasador bool   `json:"-"`
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
