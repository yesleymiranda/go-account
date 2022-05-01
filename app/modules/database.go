package modules

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/yesleymiranda/go-account/pkg/database"
)

var Database = fx.Provide(newDatabase)

func newDatabase() *gorm.DB {
	db := database.New()
	database.RunMigrations(db, nil)
	return db
}
