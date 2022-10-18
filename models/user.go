package models

type User struct {
	Name             string `json:"name,omitempty" bson:"name,omitempty"`
	Email            string `json:"email,omitempty" bson:"email,omitempty"`
	CodeforcesHandle string `json:"cfhandle,omitempty" bson:"cfhandle,omitempty"`
	Password         string `json:"password,omitempty" bson:"password,omitempty"`
}
