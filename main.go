package main

import (
	"os"

	"github.com/Heads-and-Hands/handh-school-back/config"
	"github.com/Heads-and-Hands/handh-school-back/handlers"
	"github.com/Heads-and-Hands/handh-school-back/myAdminConf"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/Heads-and-Hands/handh-school-back/bindatafs"
	"github.com/gin-gonic/gin"
)

func main() {
	if _, debug := os.LookupEnv("DEBUG"); debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	guest := router.Group("/")
	guest.GET("/", handlers.GetHandler)
	guest.POST("/", handlers.PostHandler)

	m := myAdminConf.InitAdmin()
	adminCredentials := make(gin.Accounts)
	adminCredentials[config.User] = config.Password
	admin := router.Group("/admin")
	admin.Use(gin.BasicAuth(adminCredentials))
	admin.Any("", gin.WrapF(m.ServeHTTP))
	admin.Any("/*any", gin.WrapF(m.ServeHTTP))

	router.Run()
}
