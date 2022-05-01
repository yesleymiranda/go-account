package modules

import (
	"github.com/yesleymiranda/go-toolkit/application"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/yesleymiranda/go-account/src/users"
)

var WireUp = fx.Invoke(newWireUp)

func newWireUp(app *application.App, db *gorm.DB) {
	users.Bind(app, db)
}
