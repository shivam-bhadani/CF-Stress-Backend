package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *Application) StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ticketIDString := params["ticketID"]
	ticketID, err := strconv.Atoi(ticketIDString)
	if err != nil {
		errorMessage := Error{
			Message: "Invalid Ticket ID",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	ticket, err := app.TicketStore.Query(ticketID)
	if err != nil {
		errorMessage := Error{
			Message: "Invalid Ticket ID",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	json.NewEncoder(w).Encode(*ticket)
}
