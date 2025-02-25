package model

import "github.com/google/uuid"

type Warehouse struct {
	ID      uuid.UUID `json:"id"`
	Address string    `json:"address"`
}
