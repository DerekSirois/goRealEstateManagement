package controller

import (
	"GoRealEstateManagement/auth"
	"GoRealEstateManagement/model"
	"GoRealEstateManagement/utils"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uJson := &model.UserJson{}
		err := json.NewDecoder(r.Body).Decode(uJson)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		u := &model.User{}
		err = u.GetByUsername(uJson.Username, db)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		if !model.CheckPasswordHash(uJson.Password, u.Password) {
			utils.Respond(w, r, &utils.Response{Msg: "Wrong password or username"}, http.StatusInternalServerError)
			return
		}
		token, err := auth.CreateJWTToken(int(u.ID), u.Username)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.ResponseToken{Token: token}, http.StatusOK)
	}
}

func Register(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uJson := &model.UserJson{}
		err := json.NewDecoder(r.Body).Decode(uJson)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		u := uJson.MapToUser()
		err = u.Create(db)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "User created successfully"}, http.StatusOK)
	}
}
