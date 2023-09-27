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

func GetAllPropertyByOwner(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		properties, err := model.GetAllPropertyByOwner(db, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, properties, http.StatusOK)
	}
}

func GetPropertyById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		p := &model.Property{}
		err = p.GetPropertyById(db, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, p, http.StatusOK)
	}
}

func CreateProperty(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := &model.Property{}
		err := json.NewDecoder(r.Body).Decode(p)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		err = p.CreateProperty(db)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "Property created successfully"}, http.StatusOK)
	}
}

func UpdateProperty(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := &model.Property{}
		err := json.NewDecoder(r.Body).Decode(p)
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

		err = p.UpdateProperty(db, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "Property updated successfully"}, http.StatusOK)
	}
}

func DeleteProperty(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		err = model.DeleteProperty(db, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "Property deleted successfully"}, http.StatusOK)
	}
}
