package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/romullka/MEDIA-PROJECT/internal/model"
)

type WarehouseRepository struct {
	db *pgx.Conn
}

func NewWarehouseRepository(db *pgx.Conn) *WarehouseRepository {
	return &WarehouseRepository{db: db}
}

func (r *WarehouseRepository) CreateWarehouse(ctx context.Context, address string) (model.Warehouse, error) {
	id := uuid.New()
	_, err := r.db.Exec(ctx, "INSERT INTO warehouses (id, address) VALUES (/$1, /$2)", id, address)
	if err != nil {
		return model.Warehouse{}, err
	}
	return model.Warehouse{ID: id, Address: address}, nil
}

func (r *WarehouseRepository) GetWarehouses(ctx context.Context) ([]model.Warehouse, error) {
	rows, err := r.db.Query(ctx, "SELECT id, address FROM warehouses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []model.Warehouse
	for rows.Next() {
		var warehouse model.Warehouse
		if err := rows.Scan(&warehouse.ID, &warehouse.Address); err != nil {
			return nil, err
		}
		warehouses = append(warehouses, warehouse)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return warehouses, nil
}

func (r *WarehouseRepository) GetWarehouseByID(ctx context.Context, id uuid.UUID) (model.Warehouse, error) {
	var warehouse model.Warehouse
	err := r.db.QueryRow(ctx, "SELECT id, address FROM warehouses WHERE id = /$1", id).Scan(&warehouse.ID, &warehouse.Address)
	if err != nil {
		return model.Warehouse{}, err
	}
	return warehouse, nil
}

func (r *WarehouseRepository) UpdateWarehouse(ctx context.Context, id uuid.UUID, newAddress string) (model.Warehouse, error) {
	_, err := r.db.Exec(ctx, "UPDATE warehouses SET address = /$1 WHERE id = /$2", newAddress, id)
	if err != nil {
		return model.Warehouse{}, err
	}
	return model.Warehouse{ID: id, Address: newAddress}, nil
}

func (r *WarehouseRepository) DeleteWarehouse(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, "DELETE FROM warehouses WHERE id = /$1", id)
	return err
}
