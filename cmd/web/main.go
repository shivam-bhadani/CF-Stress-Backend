package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/shivam-bhadani/cf-stress-backend/pkg/store/mongodb"
	"github.com/shivam-bhadani/cf-stress-backend/pkg/web"
)

func main() {

	ticketStore, counter, err := mongodb.NewMongoStore()
	if err != nil {
		fmt.Println(err)
	}

	app, r := web.CreateWebServer(counter, ticketStore)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	cred := handlers.AllowCredentials()
	fmt.Println(app.Counter)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins, cred)(r)))

}
