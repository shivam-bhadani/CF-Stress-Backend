package models

type Submission struct {
	ID               int    `json:"id,omitempty" bson:"id,omitempty"`
	CodeforcesHandle string `json:"cfhandle,omitempty" bson:"cfhandle,omitempty"`
	Lang             string `json:"lang,omitempty" bson:"lang,omitempty"`
}
