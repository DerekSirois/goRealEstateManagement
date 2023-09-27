package app

import (
	"GoRealEstateManagement/auth"
	"GoRealEstateManagement/controller"
	"fmt"
	"net/http"
)

func (a *App) Routes() {
	a.Router.HandleFunc("/", auth.VerifyJWT(a.index()))
	a.Router.HandleFunc("/login", controller.Login(a.DB)).Methods("POST")
	a.Router.HandleFunc("/register", controller.Register(a.DB)).Methods("POST")
	a.propertyRoutes()
}

func (a *App) propertyRoutes() {
	a.Router.HandleFunc("/api/property/owner/{id:[0-9]+}", controller.GetAllPropertyByOwner(a.DB)).Methods("GET")
	a.Router.HandleFunc("/api/property/{id:[0-9]+}", controller.GetPropertyById(a.DB)).Methods("GET")
	a.Router.HandleFunc("/api/property", controller.CreateProperty(a.DB)).Methods("POST")
	a.Router.HandleFunc("/api/property/{id:[0-9]+}", controller.UpdateProperty(a.DB)).Methods("PUT")
	a.Router.HandleFunc("/api/property/{id:[0-9]+}", controller.DeleteProperty(a.DB)).Methods("DELETE")
}

func (a *App) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Welcome to the Go real estate management tool")
	}
}
