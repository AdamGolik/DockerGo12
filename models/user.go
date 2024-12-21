package models

type User struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `json:"name" validate:"required"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Email     string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Username  string `json:"username" gorm:"not null" validate:"required"`
	Password  string `json:"password" gorm:"not null" validate:"required,min=6"`
	Telephone string `json:"telephone"`
}
