package myAdminConf

import (
	"net/http"

	"github.com/Heads-and-Hands/handh-school-back/bindatafs"
	"github.com/Heads-and-Hands/handh-school-back/models"

	"github.com/Heads-and-Hands/handh-school-back/database"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/roles"
)

func InitAdmin() *http.ServeMux {
	DB := database.CurrentDb
	DB.AutoMigrate(&models.CreateUserBody{})

	Admin := admin.New(&admin.AdminConfig{DB: DB})
	Admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))
	//bindatafs.AssetFS.Compile()

	reqRes := Admin.AddResource(&models.CreateUserBody{}, &admin.Config{
		Name: "Студенты",
		Menu: []string{"Школа разработчиков"},
	})

	reqRes.Meta(&admin.Meta{Name: "Name", Valuer: func(record interface{}, context *qor.Context) interface{} {
		if p, ok := record.(*models.CreateUserBody); ok {
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
		Name:       "Direction",
		Label:      "Направление",
		Permission: roles.Deny(roles.Delete, roles.Anyone).Deny(roles.Create, roles.Anyone).Deny(roles.Update, roles.Anyone).Allow(roles.Read, roles.Anyone),
		Config: &admin.SelectOneConfig{Collection: []string{
			"iOS", "Android", "Frontend", "Backend (Java)", "QA",
		}}})

	reqRes.IndexAttrs("-Id", "-Surname", "-Education", "-Why", "-Link")
	reqRes.EditAttrs("-Id", "-Surname")
	reqRes.Scope(&admin.Scope{Name: "iOS", Group: "Направление", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("Direction like ?", "iOS")
	}})
	reqRes.Scope(&admin.Scope{Name: "Android", Group: "Направление", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("Direction like ?", "Android")
	}})
	reqRes.Scope(&admin.Scope{Name: "Frontend", Group: "Направление", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("Direction like ?", "Frontend")
	}})
	reqRes.Scope(&admin.Scope{Name: "Backend (Java)", Group: "Направление", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("Direction like ?", "Backend (Java)")
	}})
	reqRes.Scope(&admin.Scope{Name: "QA", Group: "Направление", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("Direction like ?", "QA")
	}})

	m := http.NewServeMux()
	Admin.MountTo("/admin", m)
	return m
}
