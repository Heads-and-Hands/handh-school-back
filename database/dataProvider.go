package database

import (
	"log"
	"time"

	"github.com/Heads-and-Hands/handh-school-back/models"

	"github.com/Heads-and-Hands/handh-school-back/config"
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

func (dbp ormProvider) CreateRequest(r models.Request) {
	dbp.db.Create(&r)
}

func (dbp ormProvider) GetRequests() []models.Request {
	var requests []models.Request
	dbp.db.Table(models.Request{}.TableName()).Find(&requests)
	return requests
}

func (dbp ormProvider) UpdateRequest(p *models.Request) {
	dbp.db.Save(p)
}
