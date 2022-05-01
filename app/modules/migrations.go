package modules

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/yesleymiranda/go-account/pkg/database"
	"github.com/yesleymiranda/go-account/src/users"
)

var Migrations = fx.Invoke(newMigrations)

func newMigrations(db *gorm.DB) {
	models := declareModels()
	database.RunMigrations(db, models)
}

// DeclareModels declare all models for auto migrations
func declareModels() []interface{} {
	models := make([]interface{}, 0)
	models = append(models, &users.User{})
	return models
}
