package users

import (
	"gorm.io/gorm"

	"github.com/yesleymiranda/go-toolkit/application"
)

func Bind(app *application.App, db *gorm.DB) {
	rep := NewRepository(db)
	svc := NewService(rep)
	routes(app, svc)
}

func routes(app *application.App, svc Service) {
	app.Router.Get("/users/{id:[0-9]{1,12}}", GetByIDHandler(svc))
}
