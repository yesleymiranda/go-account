package modules

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/yesleymiranda/go-account/src/users"
	"github.com/yesleymiranda/go-toolkit/webapplication"
)

var WireUp = fx.Invoke(newWireUp)

func newWireUp(app *webapplication.App, db *gorm.DB) {
	users.Bind(app, db)
}
