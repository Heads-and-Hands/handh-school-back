package myAdminConf

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"handh-school-back/models"
	"handh-school-back/database"
	"github.com/qor/admin"
)

func InitAdmin() *http.ServeMux {
	DB := database.CurrentDb
	DB.AutoMigrate(&models.Request{})

	Admin := admin.New(&admin.AdminConfig{DB: DB})
	Admin.AddResource(&models.Request{})

	&

	m := http.NewServeMux()
	Admin.MountTo("/admin", m)
	return m
}
