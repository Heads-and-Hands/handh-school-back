package myAdminConf

import (
	"github.com/qor/roles"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"handh-school-back/models"
	"handh-school-back/database"
	"github.com/qor/admin"
	"github.com/jinzhu/gorm"
	"github.com/qor/qor"
)

func InitAdmin() *http.ServeMux {
	DB := database.CurrentDb
	DB.AutoMigrate(&models.Request{})

	Admin := admin.New(&admin.AdminConfig{DB: DB})
	reqRes := Admin.AddResource(&models.Request{}, &admin.Config{
		Name:       "Новосибирск",
		Menu:       []string{"Школа разработчиков"},
	})

	reqRes.Meta(&admin.Meta{Name: "Name", Valuer: func(record interface{}, context *qor.Context) interface{} {
		if p, ok := record.(*models.Request); ok {
			return p.Name + " " + p.Surname
		}
		return ""
	}})
	reqRes.Meta(&admin.Meta{Name: "Name", Label: "Имя"})
	reqRes.Meta(&admin.Meta{Name: "Surname", Label: "Фамилия"})
	reqRes.Meta(&admin.Meta{Name: "Education", Label: "Образование"})
	reqRes.Meta(&admin.Meta{Name: "Email", Label: "Почта"})
	reqRes.Meta(&admin.Meta{Name: "Why", Label: "Причина поступления"})
	reqRes.Meta(&admin.Meta{Name: "Link", Label: "Тестовое"})
	reqRes.Meta(&admin.Meta{
		Name: "Direction",
		Label: "Направление",
		Permission: roles.Deny(roles.Delete, roles.Anyone).Deny(roles.Create, roles.Anyone).Deny(roles.Update, roles.Anyone).Allow(roles.Read, roles.Anyone),
		Config: &admin.SelectOneConfig{Collection: []string{
			"iOS", "Android", "Frontend", "Backend (Java)",
		}}})

	reqRes.IndexAttrs("-Id", "-Surname", "-Education", "-Why", "-Link")
	reqRes.EditAttrs("-Id", "-Surname")
	reqRes.Scope(&admin.Scope{Name: "iOS", Group:"Направление", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("Direction like ?", "iOS")
	}})
	reqRes.Scope(&admin.Scope{Name: "Android", Group:"Направление", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("Direction like ?", "Android")
	}})
	reqRes.Scope(&admin.Scope{Name: "Frontend", Group:"Направление", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("Direction like ?", "Frontend")
	}})
	reqRes.Scope(&admin.Scope{Name: "Backend (Java)", Group:"Направление", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("Direction like ?", "Backend (Java)")
	}})

	m := http.NewServeMux()
	Admin.MountTo("/admin", m)
	return m
}
