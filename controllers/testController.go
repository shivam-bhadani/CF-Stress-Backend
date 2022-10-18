package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shivam-bhadani/cf-stress-backend/db"
	"github.com/shivam-bhadani/cf-stress-backend/models"
	"github.com/shivam-bhadani/cf-stress-backend/pkg/stress"
	"go.mongodb.org/mongo-driver/bson"
)

type SubmissionDetail struct {
	SubmissionID     string `json:"submission_id,omitempty"`
	CodeforcesHandle string `json:"cfhandle,omitempty"`
}

type Error struct {
	Message string `json:"error,omitempty"`
}

func (app *Application) TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	contestIDString := params["contestID"]
	problemIndex := params["problemIndex"]
	var submissionDetail SubmissionDetail
	contestID, err := strconv.Atoi(contestIDString)
	if err != nil {
		errorMessage := Error{
			Message: "ContestID must be Integer",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	json.NewDecoder(r.Body).Decode(&submissionDetail)
	problem := models.Problem{
		ContestID: contestID,
		Index:     problemIndex,
	}
	submissionID, err := strconv.Atoi(submissionDetail.SubmissionID)
	if err != nil {
		errorMessage := Error{
			Message: "SubmissionID must be Integer",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	submission := models.Submission{
		ID:               submissionID,
		CodeforcesHandle: submissionDetail.CodeforcesHandle,
		Lang:             "cpp",
	}

	app.Lock()
	app.Counter += 1
	app.Unlock()
	update_counter(app.Counter)
	ticketID := app.Counter

	go func(counter int) {
		ticket := models.Ticket{
			TicketID:   counter,
			Type:       "stress",
			Progress:   "queue",
			Problem:    problem,
			Submission: submission,
		}
		app.TicketStore.Add(&ticket)
		// processing the submission
		_ = stress.ProcessTicket(&ticket)
		app.Channel <- true
	}(app.Counter)

	go func() {
		_ = <-app.Channel
	}()

	json.NewEncoder(w).Encode(ticketID)
}

func update_counter(cnt int) error {
	client, err := db.DbConnection()
	countersCollection := client.Database("cfstress").Collection("counters")
	var counter models.Counter
	err = countersCollection.FindOne(context.TODO(), bson.M{"type": "ticket_counter"}).Decode(&counter)
	if counter.Counter > cnt {
		cnt = counter.Counter
	}
	_, err = countersCollection.UpdateOne(context.TODO(), bson.M{"type": "ticket_counter"}, bson.M{"$set": bson.M{"counter": cnt}})
	return err
}
