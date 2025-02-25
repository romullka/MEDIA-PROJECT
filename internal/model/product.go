package model

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID         `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Specs       map[string]string `json:"specs"`
	Weight      float64           `json:"weight"`
	Barcode     string            `json:"barcode"`
}
