package stress

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/shivam-bhadani/cf-stress-backend/db"
	"github.com/shivam-bhadani/cf-stress-backend/models"
	"github.com/shivam-bhadani/cf-stress-backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StressTest(ticket *models.Ticket) error {
	// getting path of required files
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	contestID := ticket.Problem.ContestID
	index := ticket.Problem.Index
	problemIndex := strings.ToLower(index)

	testPathSuffix := `/playground/contest/` + strconv.Itoa(contestID) + `/` + problemIndex
	testPath := filepath.Join(wd, testPathSuffix)
	testFileNamePrefix := "sub_id-" + strconv.Itoa(ticket.Submission.ID) + "-ticket-" + strconv.Itoa(ticket.TicketID)
	codeFileName := testFileNamePrefix + "-code"
	inputFileName := testFileNamePrefix + "-input.txt"
	participantOutputFileName := testFileNamePrefix + "-output-participant.txt"
	juryOutputFileName := testFileNamePrefix + "-output-jury.txt"
	participantSolutionPath := filepath.Join(testPath, codeFileName)
	inputFilePath := filepath.Join(testPath, inputFileName)
	participantOutputFilePath := filepath.Join(testPath, participantOutputFileName)
	juryOutputFilePath := filepath.Join(testPath, juryOutputFileName)

	binDirectoryPath := filepath.Join(wd, "bin")
	checkerScriptPath := filepath.Join(binDirectoryPath, "bash/checker.sh")
	generatorPath := filepath.Join(binDirectoryPath, "contest/"+strconv.Itoa(ticket.Problem.ContestID)+"/"+ticket.Problem.Index+"/generator")
	jurySolutionPath := filepath.Join(binDirectoryPath, "contest/"+strconv.Itoa(ticket.Problem.ContestID)+"/"+ticket.Problem.Index+"/solution")

	// arguments to pass in bash script
	args := []string{
		fmt.Sprintf("%s.cpp", participantSolutionPath),
		fmt.Sprintf("%s.cpp", jurySolutionPath),
		fmt.Sprintf("%s.cpp", generatorPath),
		inputFilePath,
		participantOutputFilePath,
		juryOutputFilePath,
		participantSolutionPath,
		jurySolutionPath,
		generatorPath,
	}

	// printing all paths
	// fmt.Println(fmt.Sprintf("%s.cpp", participantSolutionPath))
	// fmt.Println(fmt.Sprintf("%s.cpp", jurySolutionPath))
	// fmt.Println(fmt.Sprintf("%s.cpp", generatorPath))
	// fmt.Println(inputFilePath)
	// fmt.Println(participantOutputFilePath)
	// fmt.Println(juryOutputFilePath)
	// fmt.Println(participantSolutionPath)
	// fmt.Println(jurySolutionPath)
	// fmt.Println(generatorPath)

	// duration for which stress test run
	duration := 1 * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	cmd := exec.CommandContext(ctx, checkerScriptPath, args...)
	_, err = cmd.CombinedOutput()
	if err != nil {
		errorMessage := "Deadline of Stress Testing Exceeded and Couldn't Find any Wrong Testcase..."
		// fmt.Println("error in running script")
		// fmt.Println(err)
		// os.Stdout.Write(result)
		UpdateTicketError(ticket, errorMessage)
		CleanResources(ticket)
		return errors.New(errorMessage)
	}
	if ctx.Err() == context.DeadlineExceeded {
		errorMessage := "Deadline of Stress Testing Exceeded and Couldn't Find any Wrong Testcase..."
		// fmt.Println("error : deadline exceeded")
		UpdateTicketError(ticket, errorMessage)
		CleanResources(ticket)
		return errors.New(errorMessage)
	}
	input, err := utils.ReadFile(inputFilePath)
	if err != nil {
		errorMessage := "Error received in stress testing..."
		// fmt.Println("error : input file read problem")
		fmt.Println(err)
		UpdateTicketError(ticket, errorMessage)
		CleanResources(ticket)
		return errors.New(errorMessage)
	}
	participant_output, err := utils.ReadFile(participantOutputFilePath)
	if err != nil {
		errorMessage := "Error received in stress testing..."
		// fmt.Println("participant_file output file read problem")
		UpdateTicketError(ticket, errorMessage)
		CleanResources(ticket)
		return errors.New(errorMessage)
	}
	jury_output, err := utils.ReadFile(juryOutputFilePath)
	if err != nil {
		errorMessage := "Error received in stress testing..."
		// fmt.Println("error : jury_file output file read problem")
		UpdateTicketError(ticket, errorMessage)
		CleanResources(ticket)
		return errors.New(errorMessage)
	}
	err = UpdateTicketProcessed(ticket, input, participant_output, jury_output)
	if err != nil {
		CleanResources(ticket)
		return err
	}
	CleanResources(ticket)
	return nil
}

func UpdateTicketProcessed(ticket *models.Ticket, input string, participant_output string, jury_output string) error {
	client, err := db.DbConnection()
	if err != nil {
		return err
	}
	ticketsCollection := client.Database("cfstress").Collection("tickets")
	fmt.Println(participant_output)
	testCase := models.Testcase{
		Input:             input,
		JuryOutput:        jury_output,
		ParticipantOutput: participant_output,
	}
	filter := bson.D{primitive.E{Key: "ticket_id", Value: ticket.TicketID}}
	update := bson.M{"$set": bson.M{
		"progress": "processed",
		"testcase": testCase,
		"verdict":  true,
	}}
	_, err = ticketsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		// fmt.Println("error : database updation problem")
		return errors.New("Error received in stress testing...")
	}
	return nil
}
