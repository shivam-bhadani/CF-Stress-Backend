package store

import "github.com/shivam-bhadani/cf-stress-backend/models"

type TicketStore interface {
	Add(ticket *models.Ticket) error
	Query(id int) (*models.Ticket, error)
	Update(id int, updatedTicket *models.Ticket) error
	Close() error
}
