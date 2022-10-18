package controllers

import (
	"encoding/json"
	"net/http"
	"time"
)

func (app *Application) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	http.SetCookie(w,
		&http.Cookie{
			Name:    "cfstressjwt",
			Value:   "",
			Path:    "/",
			Expires: time.Unix(0, 0),
		})
	json.NewEncoder(w).Encode("success")
}
