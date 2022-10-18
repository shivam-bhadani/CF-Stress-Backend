package scraper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetExactSubmission(contestID int, submissionID int) (string, string, error) {
	url := fmt.Sprintf("https://codeforces.com/contest/%d/submission/%d", contestID, submissionID)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", "", err
	}
	verdictAccepted := doc.Find(".verdict-accepted").Text()
	code := doc.Find("#program-source-text").Text()
	return code, verdictAccepted, nil
}
