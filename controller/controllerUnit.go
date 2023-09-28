package controller

import (
	"GoRealEstateManagement/model"
	"GoRealEstateManagement/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetAllUnitByProperty(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		properties, err := model.GetAllUnitByProperty(db, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, properties, http.StatusOK)
	}
}

func GetUnitById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		u := &model.Unit{}
		err = u.GetUnitById(db, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, u, http.StatusOK)
	}
}

func CreateUnit(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &model.Unit{}
		err := json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		err = u.CreateUnit(db)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "Property created successfully"}, http.StatusOK)
	}
}

func UpdateUnit(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &model.Unit{}
		err := json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}

		err = u.UpdateUnit(db, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "Property updated successfully"}, http.StatusOK)
	}
}

func DeleteUnit(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		err = model.DeleteUnit(db, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "Property deleted successfully"}, http.StatusOK)
	}
}
