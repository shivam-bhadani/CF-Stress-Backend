package stress

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/shivam-bhadani/cf-stress-backend/models"
)

func CleanResources(ticket *models.Ticket) error {
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

	os.Remove(participantSolutionPath)
	os.Remove(fmt.Sprintf("%s.cpp", participantSolutionPath))
	os.Remove(inputFilePath)
	os.Remove(participantOutputFilePath)
	os.Remove(juryOutputFilePath)
	return nil
}
