package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/romullka/MEDIA-PROJECT/internal/service"
)

type AnalyticsHandler struct {
	analyticsService *service.AnalyticsService
}

func NewAnalyticsHandler(analyticsService *service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{analyticsService: analyticsService}
}

func (h *AnalyticsHandler) RecordSale(w http.ResponseWriter, r *http.Request) {
	var req struct {
		WarehouseID uuid.UUID `json:"warehouse_id"`
		ProductID   uuid.UUID `json:"product_id"`
		Quantity    int       `json:"quantity"`
		Price       float64   `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.analyticsService.RecordSale(r.Context(), req.WarehouseID, req.ProductID, req.Quantity, req.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sale recorded successfully"})
}

func (h *AnalyticsHandler) GetAnalyticsByWarehouse(w http.ResponseWriter, r *http.Request) {
	warehouseID := r.URL.Query().Get("warehouse_id")
	warehouseUUID, err := uuid.Parse(warehouseID)
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	analytics, err := h.analyticsService.GetAnalyticsByWarehouse(r.Context(), warehouseUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(analytics)
}

func (h *AnalyticsHandler) GetTopWarehouses(w http.ResponseWriter, r *http.Request) {
	topWarehouses, err := h.analyticsService.GetTopWarehouses(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(topWarehouses)
}
