package app

import (
	"GoRealEstateManagement/auth"
	"fmt"
	"net/http"
)

func (a *App) Routes() {
	a.Router.HandleFunc("/", auth.VerifyJWT(a.index()))
	a.Router.HandleFunc("/login", a.login()).Methods("POST")
	a.Router.HandleFunc("/register", a.register()).Methods("POST")
}

func (a *App) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Welcome to the Go real estate management tool")
	}
}
