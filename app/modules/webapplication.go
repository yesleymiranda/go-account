package modules

import (
	"go.uber.org/fx"

	"github.com/yesleymiranda/go-account/pkg/constants"
	"github.com/yesleymiranda/go-toolkit/webapplication"
)

var App = fx.Provide(newWebApplication)

func newWebApplication() *webapplication.App {
	app := webapplication.New(&webapplication.ApplicationConfig{
		Port:     constants.Port,
		WithPing: true,
	})
	app.Initialize()
	return app
}

func StartHTTPServer(app *webapplication.App) error {
	return app.ListenAndServe()
}
