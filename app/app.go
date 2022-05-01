package app

import (
	"github.com/yesleymiranda/go-account/pkg/database"
	"github.com/yesleymiranda/go-account/src/users"
	"github.com/yesleymiranda/go-toolkit/webapplication"
	"gorm.io/gorm"
)

const port = "8090"

func Run() error {
	app := webapplication.New(&webapplication.ApplicationConfig{
		Port:     port,
		WithPing: true,
	})
	app.Initialize()
	db := declareDatabase()
	wireUp(app, db)
	return app.ListenAndServe()
}

func wireUp(app *webapplication.App, db *gorm.DB) {
	users.Bind(app, db)
}

func declareDatabase() *gorm.DB {
	db := database.New()
	database.RunMigrations(db)
	return db
}
