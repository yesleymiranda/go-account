package app

import (
	"go.uber.org/fx"

	"github.com/yesleymiranda/go-account/app/modules"
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
