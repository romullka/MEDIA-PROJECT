package handler

import (
	"encoding/json"
	"net/http"

	"github.com/romullka/MEDIA-PROJECT/internal/repository"
)

type WarehouseHandler struct {
	repo *repository.WarehouseRepository
}

func NewWarehouseHandler(repo *repository.WarehouseRepository) *WarehouseHandler {
	return &WarehouseHandler{repo: repo}
}

func (h *WarehouseHandler) CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouse struct {
		Address string `json:"address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newWarehouse, err := h.repo.CreateWarehouse(r.Context(), warehouse.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newWarehouse)
}

func (h *WarehouseHandler) GetWarehouses(w http.ResponseWriter, r *http.Request) {
	warehouses, err := h.repo.GetWarehouses(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(warehouses)
}
