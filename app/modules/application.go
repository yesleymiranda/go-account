package modules

import (
	"go.uber.org/fx"

	"github.com/yesleymiranda/go-account/pkg/constants"
	"github.com/yesleymiranda/go-toolkit/application"
)

var App = fx.Provide(newWebApplication)

func newWebApplication() *application.App {
	app := application.New(&application.Config{
		Port:     constants.Port,
		WithPing: true,
	})
	app.Initialize()
	return app
}

func StartHTTPServer(app *application.App) error {
	return app.ListenAndServe()
}
