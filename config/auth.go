package config

import "os"

var User string
var Password string

func init() {
	User = os.Getenv("AUTH_USER")
	Password = os.Getenv("AUTH_PASSWORD")
}
