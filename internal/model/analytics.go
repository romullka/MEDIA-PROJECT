package model

import (
	"github.com/google/uuid"
)

type Analytics struct {
	WarehouseID  uuid.UUID
	ProductID    uuid.UUID
	SoldQuantity int
	TotalSum     float64
}

type AnalyticsData struct {
	TotalQuantity int     `json:"total_quantity"`
	TotalSum      float64 `json:"total_sum"`
}
