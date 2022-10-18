package stress

import (
	"context"
	"errors"
	"strings"

	"github.com/shivam-bhadani/cf-stress-backend/db"
	"github.com/shivam-bhadani/cf-stress-backend/models"
	"github.com/shivam-bhadani/cf-stress-backend/pkg/scraper"
	"go.mongodb.org/mongo-driver/bson"
)

func ProcessTicket(ticket *models.Ticket) error {
	// ticketID := ticket.TicketID
	contestID := ticket.Problem.ContestID
	index := ticket.Problem.Index
	problemIndex := strings.ToLower(index)
	client, err := db.DbConnection()
	if err != nil {
		return err
	}
	problemsCollection := client.Database("cfstress").Collection("problems")

	submissionID := ticket.Submission.ID

	code, verdictAccepted, err := scraper.GetExactSubmission(contestID, submissionID)
	// fmt.Println(code)
	// if err != nil {
	// 	return err
	// }
	if verdictAccepted == "Accepted" {
		errorMessage := "This Problem is already accepted in CodeForces, therefore no counter test case available"
		_ = UpdateTicketError(ticket, errorMessage)

		return errors.New(errorMessage)
	}

	filter := bson.M{
		"contest_id": contestID,
		"index":      problemIndex,
	}
	res := problemsCollection.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		errorMessage := "Currently stress test for this problem is not available. We will surely add this problem after some days."
		_ = UpdateTicketError(ticket, errorMessage)
		return errors.New(errorMessage)
	}

	err = MakeFilesForStressTest(ticket, code)
	if err != nil {
		return err
	}

	err = StressTest(ticket)
	if err != nil {
		return err
	}

	return nil
}
