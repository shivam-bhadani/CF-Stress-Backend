package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shivam-bhadani/cf-stress-backend/db"
	"github.com/shivam-bhadani/cf-stress-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (app *Application) SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	client, err := db.DbConnection()
	userCollection := client.Database("cfstress").Collection("users")
	res := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email})
	if res.Err() == nil {
		json.NewEncoder(w).Encode("This email already exists")
		fmt.Println(err)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		json.NewEncoder(w).Encode("Password does not match")
		return
	}
	hashed := string(hash)
	user.Password = hashed
	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode("success")
}
