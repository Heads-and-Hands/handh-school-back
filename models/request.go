package models

type Request struct {
	Id 			int    	`json:"id"`
	Name 		string 	`json:"first_name"`
	Surname		string 	`json:"last_name"`
	Education 	string 	`json:"education"`
	Why			string 	`json:"reason"`
	Direction	string  `json:"course"`
}

func (Request) TableName() string {
	return "requests"
}
