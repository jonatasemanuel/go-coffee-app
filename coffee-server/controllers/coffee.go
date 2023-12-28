package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jonatasemanuel/coffee-server/helpers"
	"github.com/jonatasemanuel/coffee-server/services"
)

var coffee services.Coffee

// GET/coffees
func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	all, err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"coffees": all})
}

func CreateCoffe(w http.ResponseWriter, r *http.Request) {
	var coffeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	coffeCreated, err := coffee.CreateCoffee(coffeData)
	// check
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, coffeCreated)
}
