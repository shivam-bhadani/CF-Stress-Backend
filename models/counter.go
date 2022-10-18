package models

type Counter struct {
	Type    string `json:"type,omitempty" bson:"type,omitempty"`
	Counter int    `json:"counter,omitempty" bson:"counter,omitempty"`
}
