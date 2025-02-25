package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/romullka/MEDIA-PROJECT/internal/service"
)

type InventoryHandler struct {
	inventoryService *service.InventoryService
}

func NewInventoryHandler(inventoryService *service.InventoryService) *InventoryHandler {
	return &InventoryHandler{inventoryService: inventoryService}
}

func (h *InventoryHandler) CreateInventory(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ProductID   uuid.UUID `json:"product_id"`
		WarehouseID uuid.UUID `json:"warehouse_id"`
		Quantity    int       `json:"quantity"`
		Price       float64   `json:"price"`
		Discount    float64   `json:"discount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.inventoryService.CreateInventory(r.Context(), input.ProductID, input.WarehouseID, input.Quantity, input.Price, input.Discount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *InventoryHandler) PurchaseProducts(w http.ResponseWriter, r *http.Request) {
	var input struct {
		WarehouseID uuid.UUID `json:"warehouse_id"`
		ProductID   uuid.UUID `json:"product_id"`
		Quantity    int       `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	totalPrice, err := h.inventoryService.PurchaseProducts(r.Context(), input.WarehouseID, input.ProductID, input.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Purchase successful", "total_price": totalPrice})
}

func (h *InventoryHandler) GetAnalyticsByWarehouse(w http.ResponseWriter, r *http.Request) {
	warehouseID := r.URL.Query().Get("warehouse_id")
	id, err := uuid.Parse(warehouseID)
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	analytics, err := h.inventoryService.GetAnalyticsByWarehouse(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(analytics)
}
