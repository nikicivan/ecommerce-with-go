package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email" gorm:"unique"`
	Password    []byte   `json:"-"`
	IsAmbasador bool     `json:"-"`
	Revenue     *float64 `json:"revenue,omitempty" gorm:"-"`
}

type Admin User

type Ambassador User

// SetPassword generates hashed password for user
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

// CompareHashedPassword compares hashedPassword and password from request
func (user *User) CompareHashedPassword(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// Name returns user full name
func (user *User) FullName() string {
	return user.FirstName + " " + user.LastName
}

// CalculateRevenue calculates the ambassador revenue
func (ambassador *Ambassador) CalculateRevenue(db *gorm.DB) {
	var orders []Order

	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   ambassador.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.AmbassadorRevenue
		}
	}

	ambassador.Revenue = &revenue
}

// CalculateRevenue calculates the admin revenue
func (admin *Admin) CalculateRevenue(db *gorm.DB) {
	var orders []Order

	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   admin.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.AdminRevenue
		}
	}

	admin.Revenue = &revenue
}
