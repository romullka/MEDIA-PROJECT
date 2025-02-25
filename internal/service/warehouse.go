package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/romullka/MEDIA-PROJECT/internal/model"
	"github.com/romullka/MEDIA-PROJECT/internal/repository"
)

type WarehouseService struct {
	repo *repository.WarehouseRepository
}

func NewWarehouseService(repo *repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{repo: repo}
}

func (s *WarehouseService) CreateWarehouse(ctx context.Context, address string) (model.Warehouse, error) {
	return s.repo.CreateWarehouse(ctx, address)
}

func (s *WarehouseService) GetWarehouses(ctx context.Context) ([]model.Warehouse, error) {
	return s.repo.GetWarehouses(ctx)
}

func (s *WarehouseService) GetWarehouseByID(ctx context.Context, id uuid.UUID) (model.Warehouse, error) {
	return s.repo.GetWarehouseByID(ctx, id)
}

func (s *WarehouseService) UpdateWarehouse(ctx context.Context, id uuid.UUID, newAddress string) (model.Warehouse, error) {
	return s.repo.UpdateWarehouse(ctx, id, newAddress)
}

func (s *WarehouseService) DeleteWarehouse(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteWarehouse(ctx, id)
}
