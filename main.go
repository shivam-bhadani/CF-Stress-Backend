package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shivam-bhadani/cf-stress-backend/controllers"
)

func main() {
	r := mux.NewRouter()
	corsObj := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"})
	r.HandleFunc("/api/signup", controllers.SignupHandler).Methods("POST")
	r.HandleFunc("/api/login", controllers.LoginHandler).Methods("POST")
	r.HandleFunc("/api/user", controllers.UserHandler).Methods("GET")
	r.HandleFunc("/api/logout", controllers.LogoutHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(r)))
}
