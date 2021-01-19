package models

import "github.com/jinzhu/gorm"

type Request struct {
	gorm.Model
	Name 		string 	`json:"first_name"`
	Surname		string 	`json:"last_name"`
	Education 	string 	`json:"education"`
	Email	 	string 	`json:"email"`
	Why			string 	`json:"reason"`
	Direction	string  `json:"course"`
	Link		string  `json:"link"`
}

func (Request) TableName() string {
	return "requests"
}
