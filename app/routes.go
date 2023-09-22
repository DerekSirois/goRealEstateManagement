package app

import (
	"fmt"
	"net/http"
)

func (a *App) Routes() {
	a.Router.HandleFunc("/", a.index())
}

func (a *App) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Welcome to the Go real estate management tool")
	}
}
