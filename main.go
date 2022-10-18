package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/shivam-bhadani/cf-stress-backend/pkg/store/mongodb"
	"github.com/shivam-bhadani/cf-stress-backend/pkg/web"
)

func main() {

	ticketStore, counter, err := mongodb.NewMongoStore()
	if err != nil {
		fmt.Println(err)
	}
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app, r := web.CreateWebServer(counter, ticketStore)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"https://cfstress.vercel.app/"})
	cred := handlers.AllowCredentials()
	fmt.Println(app.Counter)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins, cred)(r)))
}
