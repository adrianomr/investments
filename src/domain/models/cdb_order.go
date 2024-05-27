package models

import (
	"github.com/google/uuid"
	"time"
)

type CdbOrderType string

const (
	CdbOrderTypeBuy  CdbOrderType = "BUY"
	CdbOrderTypeSell CdbOrderType = "SELL"
)

type CdbOrder struct {
	ID     uuid.UUID    `json:"id"`
	UserID string       `json:"user_id"`
	Type   CdbOrderType `json:"type" validate:"required"`
	Amount float64      `json:"amount" validate:"required"`
	CdbId  uuid.UUID    `json:"cdb_id"`
	Date   time.Time    `json:"date"`
}
