package app

import (
	"GoRealEstateManagement/model"
	"GoRealEstateManagement/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a *App) getAllPropertyByOwner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		properties, err := model.GetAllPropertyByOwner(a.DB, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, properties, http.StatusOK)
	}
}

func (a *App) getPropertyById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		p := &model.Property{}
		err = p.GetById(a.DB, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, p, http.StatusOK)
	}
}

func (a *App) createProperty() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := &model.Property{}
		err := json.NewDecoder(r.Body).Decode(p)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		err = p.Create(a.DB)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "Property created successfully"}, http.StatusOK)
	}
}

func (a *App) updateProperty() http.HandlerFunc {
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

		err = p.Update(a.DB, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "Property updated successfully"}, http.StatusOK)
	}
}

func (a *App) deleteProperty() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		err = model.DeleteProperty(a.DB, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "Property deleted successfully"}, http.StatusOK)
	}
}
