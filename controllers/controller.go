package controllers

import (
	"sync"

	"github.com/shivam-bhadani/cf-stress-backend/pkg/store"
)

type Application struct {
	Counter     int
	TicketStore store.TicketStore
	Channel     chan bool
	sync.Mutex
}
