package app

import (
	"GoRealEstateManagement/app/model"
	"encoding/json"
	"net/http"
)

func (a *App) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uJson := &model.UserJson{}
		err := json.NewDecoder(r.Body).Decode(uJson)
		if err != nil {
			a.respond(w, r, &Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		u := &model.User{}
		err = u.GetByUsername(uJson.Username, a.DB)
		if err != nil {
			a.respond(w, r, &Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		if !model.CheckPasswordHash(uJson.Password, u.Password) {
			a.respond(w, r, &Response{Msg: "Wrong password or username"}, http.StatusInternalServerError)
			return
		}
		a.respond(w, r, &Response{Msg: "You are logged in"}, http.StatusOK)
	}
}

func (a *App) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uJson := &model.UserJson{}
		err := json.NewDecoder(r.Body).Decode(uJson)
		if err != nil {
			a.respond(w, r, &Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		u := uJson.MapToUser()
		err = u.Create(a.DB)
		if err != nil {
			a.respond(w, r, &Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		a.respond(w, r, &Response{Msg: "User created successfully"}, http.StatusOK)
	}
}
