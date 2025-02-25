package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/romullka/MEDIA-PROJECT/internal/model"
)

type AnalyticsRepository struct {
	db *pgx.Conn
}

func NewAnalyticsRepository(db *pgx.Conn) *AnalyticsRepository {
	return &AnalyticsRepository{db: db}
}

func (r *AnalyticsRepository) RecordSale(ctx context.Context, analytics model.Analytics) error {
	query := `INSERT INTO sales (warehouse_id, product_id, sold_quantity, total_sum) VALUES (\$1, \$2, \$3, \$4)`
	_, err := r.db.Exec(ctx, query, analytics.WarehouseID, analytics.ProductID, analytics.SoldQuantity, analytics.TotalSum)
	return err
}

func (r *AnalyticsRepository) GetAnalyticsByWarehouse(ctx context.Context, warehouseID uuid.UUID) (model.AnalyticsData, error) {
	query := `SELECT SUM(sold_quantity) as total_quantity, SUM(total_sum) as total_sum FROM sales WHERE warehouse_id = \$1`

	var data model.AnalyticsData
	err := r.db.QueryRow(ctx, query, warehouseID).Scan(&data.TotalQuantity, &data.TotalSum)
	if err != nil {
		return model.AnalyticsData{}, err
	}

	return data, nil
}

func (r *AnalyticsRepository) GetTopWarehouses(ctx context.Context) ([]model.Analytics, error) {
	query := `SELECT warehouse_id, SUM(sold_quantity) as total_quantity, SUM(total_sum) as total_sum FROM sales GROUP BY warehouse_id
    ORDER BY SUM(total_sum) DESC LIMIT 10`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topWarehouses []model.Analytics
	for rows.Next() {
		var analytics model.Analytics
		if err := rows.Scan(&analytics.WarehouseID, &analytics.SoldQuantity, &analytics.TotalSum); err != nil {
			return nil, err
		}
		topWarehouses = append(topWarehouses, analytics)
	}

	return topWarehouses, nil
}
