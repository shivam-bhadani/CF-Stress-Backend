package models

type Problem struct {
	ContestID int    `json:"contest_id,omitempty" bson:"contestId,omitempty"`
	Index     string `json:"index,omitempty" bson:"index,omitempty"`
}
