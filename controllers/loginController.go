package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/shivam-bhadani/cf-stress-backend/auth"
	"github.com/shivam-bhadani/cf-stress-backend/db"
	"github.com/shivam-bhadani/cf-stress-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var credentials auth.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	var user models.User
	client, err := db.DbConnection()
	userCollection := client.Database("cfstress").Collection("users")
	err = userCollection.FindOne(context.TODO(), bson.M{"email": credentials.Email}).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode("No user found")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		json.NewEncoder(w).Encode("Password is incorrect")
		return
	}
	token, err := auth.GenerateJWT(credentials.Email)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "cfstressjwt",
			Value:   token,
			Expires: time.Now().Add(24 * 28 * time.Hour),
			Path:    "/",
		})
	json.NewEncoder(w).Encode(user)
}
