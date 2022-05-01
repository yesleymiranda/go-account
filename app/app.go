package app

import (
	"github.com/yesleymiranda/go-account/pkg/database"
	"github.com/yesleymiranda/go-toolkit/webapplication"
	"gorm.io/gorm"
)

const port = "8080"

func Run() error {
	app := webapplication.New(&webapplication.ApplicationConfig{
		Port:     port,
		WithPing: true,
	})
	app.Initialize()
	_ = runDatabase()
	return app.ListenAndServe()
}

func runDatabase() *gorm.DB {
	db := database.New()
	database.RunMigrations(db)
	return db
}
