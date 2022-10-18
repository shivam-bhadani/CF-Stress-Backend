package models

type Ticket struct {
	TicketID   int        `json:"ticket_id,omitempty" bson:"ticket_id,omitempty"`
	Type       string     `json:"type,omitempty" bson:"type,omitempty"`
	Progress   string     `json:"progress,omitempty" bson:"progress,omitempty"`
	Verdict    bool       `json:"verdict,omitempty" bson:"verdict,omitempty"`
	Problem    Problem    `json:"problem,omitempty" bson:"problem,omitempty"`
	Submission Submission `json:"submission,omitempty" bson:"submission,omitempty"`
	Testcase   Testcase   `json:"testcase,omitempty" bson:"testcase,omitempty"`
	Error      string     `json:"error,omitempty" bson:"error,omitempty"`
}
