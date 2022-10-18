package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/shivam-bhadani/cf-stress-backend/models"
	gomail "gopkg.in/mail.v2"
)

func (app *Application) ContactController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact models.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "cfstress@gmail.com")
	m.SetHeader("To", "cfstress@gmail.com")
	m.SetHeader("Subject", contact.Subject)
	msg := fmt.Sprintf("Name : %s\nEmail: %s\n%s", contact.Name, contact.Email, contact.Message)
	m.SetBody("text/plain", msg)

	d := gomail.NewDialer("smtp.gmail.com", 587, "cfstress@gmail.com", os.Getenv("GMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode("success")
}
