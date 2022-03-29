package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConnections()
	defer db.Close()

	db.Find(&personas)
	json, _ := json.Marshal(personas)
	commons.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnections()
	defer db.Close()

	db.Find(&personas, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func Save(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnections()
	defer db.Close()

	if error != nil {
		log.Fatal(error)
		commons.SendResponse(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&persona).Error

	if error != nil {
		log.Fatal(error)
		commons.SendResponse(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)

	commons.SendResponse(writer, http.StatusCreated, json)

}

func Delete(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnections()
	defer db.Close()

	db.Find(&personas, id)

	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
