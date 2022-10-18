package mongodb

import (
	"context"

	"github.com/shivam-bhadani/cf-stress-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (store *mongoStore) Add(ticket *models.Ticket) error {
	_, err := store.ticketsCollection.InsertOne(context.TODO(), ticket)
	if err != nil {
		return err
	}
	return nil
}

func (store *mongoStore) Query(id int) (*models.Ticket, error) {
	var ticket models.Ticket
	err := store.ticketsCollection.FindOne(context.TODO(), bson.D{{"ticket_id", id}}).Decode(&ticket)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (store *mongoStore) Update(id int, updatedTicket *models.Ticket) error {
	filter := bson.D{primitive.E{Key: "ticket_id", Value: id}}
	update := bson.M{"$set": bson.M{
		"progress": updatedTicket.Progress,
		"verdict":  updatedTicket.Verdict,
		"testcase": updatedTicket.Testcase,
	}}
	_, err := store.ticketsCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (store *mongoStore) Close() error {
	return store.mongoClient.Disconnect(context.TODO())
}
