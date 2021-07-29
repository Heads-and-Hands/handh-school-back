package models

import "github.com/jinzhu/gorm"

type CreateUserBody struct {
	gorm.Model `json:"-"`
	Name       string `json:"first_name" binding:"required"`
	Surname    string `json:"last_name" binding:"required"`
	Education  string `json:"education" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Why        string `json:"reason" binding:"required"`
	Direction  string `json:"course" binding:"required"`
	Link       string `json:"link" binding:"required"`
}

func (CreateUserBody) TableName() string {
	return "requests"
}
