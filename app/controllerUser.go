package app

import (
	"GoRealEstateManagement/auth"
	"GoRealEstateManagement/model"
	"encoding/json"
	"net/http"
)

func (a *App) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uJson := &model.UserJson{}
		err := json.NewDecoder(r.Body).Decode(uJson)
		if err != nil {
			a.Respond(w, r, &Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		u := &model.User{}
		err = u.GetByUsername(uJson.Username, a.DB)
		if err != nil {
			a.Respond(w, r, &Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		if !model.CheckPasswordHash(uJson.Password, u.Password) {
			a.Respond(w, r, &Response{Msg: "Wrong password or username"}, http.StatusInternalServerError)
			return
		}
		token, err := auth.CreateJWTToken(int(u.ID), u.Username)
		if err != nil {
			a.Respond(w, r, &Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		a.Respond(w, r, &ResponseToken{Token: token}, http.StatusOK)
	}
}

func (a *App) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uJson := &model.UserJson{}
		err := json.NewDecoder(r.Body).Decode(uJson)
		if err != nil {
			a.Respond(w, r, &Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		u := uJson.MapToUser()
		err = u.Create(a.DB)
		if err != nil {
			a.Respond(w, r, &Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		a.Respond(w, r, &Response{Msg: "User created successfully"}, http.StatusOK)
	}
}
