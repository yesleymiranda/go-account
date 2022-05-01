package database

import (
	"os"

	"github.com/yesleymiranda/go-account/pkg/constants"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	if os.Getenv(constants.Platform) == constants.Prod {
		return nil
	}
	return newSqlite()
}

func newSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to exec db(): " + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	return db
}
