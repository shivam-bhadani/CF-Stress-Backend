package web

import (
	"github.com/gorilla/mux"
	"github.com/shivam-bhadani/cf-stress-backend/controllers"
	"github.com/shivam-bhadani/cf-stress-backend/pkg/store"
)

func CreateWebServer(counter int, ticketStore store.TicketStore) (app *controllers.Application, r *mux.Router) {
	bufferCapacity := 50
	app = &controllers.Application{
		Counter:     counter,
		TicketStore: ticketStore,
		Channel:     make(chan bool, bufferCapacity),
	}
	r = mux.NewRouter()

	r.HandleFunc("/api/signup", app.SignupHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/login", app.LoginHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/user", app.UserHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/logout", app.LogoutHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/contact", app.ContactController).Methods("POST", "OPTIONS")

	r.HandleFunc("/api/test/{contestID}/{problemIndex}", app.TestHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/status/{ticketID}", app.StatusHandler).Methods("GET")
	return
}
