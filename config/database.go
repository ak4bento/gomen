package config

import (
	"gomen/database"
	"gorm.io/gorm"
)

var db database.PostgreSQL

type Database struct {
}

func (db *Database) Default() {

}

func (db *Database) Connections() {

}

func init() {
	if err := db.Config.Parse(); err != nil {
		panic(err)
	}
}

func DB() (database *gorm.DB) {
	database, _ = db.Connect(&gorm.Config{})
	return database.Session(&gorm.Session{NewDB: true})
}
