package app

import (
	"github.com/yesleymiranda/go-account/pkg/database"
	"github.com/yesleymiranda/go-toolkit/webapplication"
)

const port = "8080"

func Run() error {
	app := webapplication.New(&webapplication.ApplicationConfig{
		Port:     port,
		WithPing: true,
	})
	app.Initialize()
	database.New()
	return app.ListenAndServe()
}
