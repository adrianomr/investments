package models

import "github.com/google/uuid"

type CdbOrderType string

type CdbOrder struct {
	ID            uuid.UUID    `json:"id"`
	UserID        string       `json:"user_id"`
	Type          CdbOrderType `json:"type"`
	Amount        float64      `json:"amount"`
	CdiPercentage float64      `json:"cdiPercentage"`
}

type Cdb struct {
	ID     uuid.UUID `json:"id"`
	Amount float64   `json:"amount"`
}
