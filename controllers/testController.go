package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type SubmissionDetail struct {
	SubmissionID string
}

func (app *Application) TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	contestID := params["contestID"]
	problemIndex := params["problemIndex"]
	var submissionDetail SubmissionDetail
	json.NewDecoder(r.Body).Decode(&submissionDetail)
	data := map[string]string{
		"contestID":    contestID,
		"problemIndex": problemIndex,
		"submissionID": submissionDetail.SubmissionID,
	}
	json.NewEncoder(w).Encode(data)
}
