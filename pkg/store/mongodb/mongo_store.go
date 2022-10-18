package mongodb

import (
	"context"

	"github.com/shivam-bhadani/cf-stress-backend/db"
	"github.com/shivam-bhadani/cf-stress-backend/models"
	"github.com/shivam-bhadani/cf-stress-backend/pkg/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoStore struct {
	mongoClient        *mongo.Client
	countersCollection *mongo.Collection
	ticketsCollection  *mongo.Collection
}

func NewMongoStore() (store.TicketStore, int, error) {
	client, err := db.DbConnection()
	if err != nil {
		return nil, -1, err
	}
	mStore := new(mongoStore)
	mStore.mongoClient = client
	var counter models.Counter
	mStore.countersCollection = client.Database("cfstress").Collection("counters")
	mStore.ticketsCollection = client.Database("cfstress").Collection("tickets")
	err = mStore.countersCollection.FindOne(context.TODO(), bson.M{"type": "ticket_counter"}).Decode(&counter)
	if err == nil {
		return mStore, counter.Counter, nil
	}
	counter.Type = "ticket_counter"
	counter.Counter = 1
	_, err = mStore.countersCollection.InsertOne(context.TODO(), counter)
	if err != nil {
		return nil, -1, err
	}
	return mStore, counter.Counter, nil
}
