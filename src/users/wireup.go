package users

import (
	"github.com/yesleymiranda/go-toolkit/webapplication"
	"gorm.io/gorm"
)

func Bind(app *webapplication.App, db *gorm.DB) {
	rep := NewRepository(db)
	svc := NewService(rep)
	routes(app, svc)
}

func routes(app *webapplication.App, svc Service) {
	app.Router.HandleFunc("/users/{id:[0-9]{1,12}}", MakeGetByIDHandler(svc)).Methods("GET")
}
