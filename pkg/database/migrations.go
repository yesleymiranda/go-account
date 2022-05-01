package database

import (
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	autoMigrate(db, DeclareModels())
}

func autoMigrate(db *gorm.DB, models []interface{}) {
	for _, m := range models {
		err := db.AutoMigrate(m)
		if err != nil {
			panic("error on auto migrate:" + err.Error())
		}
	}
}
