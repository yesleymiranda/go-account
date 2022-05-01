package app

import (
	"github.com/yesleymiranda/go-account/app/modules"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/yesleymiranda/go-account/src/users"
	"github.com/yesleymiranda/go-toolkit/webapplication"
)

func StartApp() *fx.App {
	return fx.New(
		fx.Options(
			modules.Database,
			modules.Migrations,
			modules.App,
			modules.WireUp,
		),
		fx.Invoke(modules.StartHTTPServer),
	)
}

func wireUp(app *webapplication.App, db *gorm.DB) {
	users.Bind(app, db)
}
