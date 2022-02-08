package models

type Link struct {
	BaseModel
	Code     string    `json:"code"`
	UserId   string    `json:"user_id"`
	User     User      `json:"user" gorm:"foreignKey:UserId"`
	Products []Product `json:"products" gorm:"many2many:link_products"`
	Orders   []Order   `json:"orders" gorm:"-"`
}
