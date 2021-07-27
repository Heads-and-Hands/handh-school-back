package database

import (
	"log"
	"time"

	"github.com/Heads-and-Hands/handh-school-back/models"

	"github.com/Heads-and-Hands/handh-school-back/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ormProvider struct {
	db *gorm.DB
}

var OrmProvider ormProvider
var CurrentDb *gorm.DB

func init() {
	time.Sleep(config.StartDelay)

	CurrentDb = getDB()
	OrmProvider = ormProvider{
		db: CurrentDb,
	}
}

func (dbp ormProvider) Close() {
	dbp.db.Close()
}

func getDB() *gorm.DB {
	dbString := config.DbString
	log.Printf("Open DB %s \n", dbString)
	newDb, err := gorm.Open("mysql", dbString)

	if err != nil {
		log.Println(err)
		return nil
	}
	return newDb
}

func (dbp ormProvider) CreateRequest(r models.CreateUserBody) {
	dbp.db.Create(&r)
}

func (dbp ormProvider) GetRequests() []models.CreateUserBody {
	var requests []models.CreateUserBody
	dbp.db.Table(models.CreateUserBody{}.TableName()).Find(&requests)
	return requests
}

func (dbp ormProvider) UpdateRequest(p *models.CreateUserBody) {
	dbp.db.Save(p)
}
