package models

import (
	"github.com/google/uuid"
	"time"
)

type Cdi struct {
	ID   uuid.UUID `json:"id"`
	Rate float64   `json:"rate"`
	Date time.Time
}
