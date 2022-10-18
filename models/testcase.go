package models

type Testcase struct {
	Input             string `json:"input,omitempty" bson:"input,omitempty"`
	JuryOutput        string `json:"jury_output,omitempty" bson:"jury_output,omitempty"`
	ParticipantOutput string `json:"participant_output,omitempty" bson:"participant_output,omitempty"`
}
