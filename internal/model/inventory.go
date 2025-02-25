package model

import (
	"github.com/google/uuid"
)

type Inventory struct {
	ID          uuid.UUID
	ProductID   uuid.UUID
	WarehouseID uuid.UUID
	Quantity    int
	Price       float64
	Discount    float64
}
