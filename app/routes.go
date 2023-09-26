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
	a.propertyRoutes()
}

func (a *App) propertyRoutes() {
	a.Router.HandleFunc("/api/property/owner/{id:[0-9]+}", a.getAllPropertyByOwner()).Methods("GET")
	a.Router.HandleFunc("/api/property/{id:[0-9]+}", a.getPropertyById()).Methods("GET")
	a.Router.HandleFunc("/api/property", a.createProperty()).Methods("POST")
	a.Router.HandleFunc("/api/property/{id:[0-9]+}", a.updateProperty()).Methods("PUT")
	a.Router.HandleFunc("/api/property/{id:[0-9]+}", a.deleteProperty()).Methods("DELETE")
}

func (a *App) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Welcome to the Go real estate management tool")
	}
}
