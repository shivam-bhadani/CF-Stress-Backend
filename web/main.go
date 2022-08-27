package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shivam-bhadani/cf-stress-backend/controllers"
)

func main() {
	app := new(controllers.Application)
	r := mux.NewRouter()
	corsObj := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"})
	r.HandleFunc("/api/signup", app.SignupHandler).Methods("POST")
	r.HandleFunc("/api/login", app.LoginHandler).Methods("POST")
	r.HandleFunc("/api/user", app.UserHandler).Methods("GET")
	r.HandleFunc("/api/logout", app.LogoutHandler).Methods("POST")
	r.HandleFunc("/api/test/{contestID}/{problemIndex}", app.TestHandler).Methods("POST")
	r.HandleFunc("/api/status/{ticketID}", app.StatusHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(r)))
}
