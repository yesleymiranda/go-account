package main

import (
	"errors"
	"net/http"

	"github.com/yesleymiranda/go-account/app"
)

func main() {
	err := app.Run()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
