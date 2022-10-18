package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/shivam-bhadani/cf-stress-backend/auth"
	"github.com/shivam-bhadani/cf-stress-backend/db"
	"github.com/shivam-bhadani/cf-stress-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *Application) UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cookie, err := r.Cookie("cfstressjwt")
	if err != nil {
		json.NewEncoder(w).Encode("unauthenticated")
		return
	}
	tokenString := cookie.Value
	err, claims := auth.ValidateToken(tokenString)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	var user models.User
	client, err := db.DbConnection()
	userCollection := client.Database("cfstress").Collection("users")
	err = userCollection.FindOne(context.TODO(), bson.M{"email": claims.Email}).Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(user)
}
