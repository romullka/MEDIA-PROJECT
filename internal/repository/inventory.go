package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/romullka/MEDIA-PROJECT/internal/model"
)

type InventoryRepository struct {
	db *pgx.Conn
}

func NewInventoryRepository(db *pgx.Conn) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) CreateInventory(ctx context.Context, productID, warehouseID uuid.UUID, quantity int, price, discount float64) (model.Inventory, error) {
	id := uuid.New()
	_, err := r.db.Exec(ctx, "INSERT INTO inventory (id, product_id, warehouse_id, quantity, price, discount) VALUES (/$1, /$2, /$3, /$4, /$5, /$6)", id, productID, warehouseID, quantity, price, discount)
	if err != nil {
		return model.Inventory{}, err
	}
	return model.Inventory{ID: id, ProductID: productID, WarehouseID: warehouseID, Quantity: quantity, Price: price, Discount: discount}, nil
}

func (r *InventoryRepository) UpdateInventoryQuantity(ctx context.Context, productID, warehouseID uuid.UUID, quantity int) error {
	_, err := r.db.Exec(ctx, "UPDATE inventory SET quantity = quantity + /$1 WHERE product_id = /$2 AND warehouse_id = /$3", quantity, productID, warehouseID)
	return err
}

func (r *InventoryRepository) CreateDiscount(ctx context.Context, productID, warehouseID uuid.UUID, discount float64) error {
	_, err := r.db.Exec(ctx, "UPDATE inventory SET discount = /$1 WHERE product_id = /$2 AND warehouse_id = /$3", discount, productID, warehouseID)
	return err
}

func (r *InventoryRepository) GetProductsByWarehouse(ctx context.Context, warehouseID uuid.UUID, limit, offset int) ([]model.Inventory, error) {
	rows, err := r.db.Query(ctx, "SELECT id, product_id, warehouse_id, quantity, price, discount FROM inventory WHERE warehouse_id = /$1 LIMIT /$2 OFFSET /$3", warehouseID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []model.Inventory
	for rows.Next() {
		var inventory model.Inventory
		if err := rows.Scan(&inventory.ID, &inventory.ProductID, &inventory.WarehouseID, &inventory.Quantity, &inventory.Price, &inventory.Discount); err != nil {
			return nil, err
		}
		inventories = append(inventories, inventory)
	}

	return inventories, nil
}

func (r *InventoryRepository) GetProductDetails(ctx context.Context, productID, warehouseID uuid.UUID) (model.Inventory, error) {
	var inventory model.Inventory
	err := r.db.QueryRow(ctx, "SELECT id, product_id, warehouse_id, quantity, price, discount FROM inventory WHERE product_id = /$1 AND warehouse_id = /$2", productID, warehouseID).Scan(&inventory.ID, &inventory.ProductID, &inventory.WarehouseID, &inventory.Quantity, &inventory.Price, &inventory.Discount)
	if err != nil {
		return model.Inventory{}, err
	}
	return inventory, nil
}
func (r *InventoryRepository) CalculateTotalPrice(ctx context.Context, warehouseID uuid.UUID, products map[uuid.UUID]int) (float64, error) {
	var totalPrice float64
	for productID, quantity := range products {
		var price, discount float64
		err := r.db.QueryRow(ctx, "SELECT price, discount FROM inventory WHERE product_id = /$1 AND warehouse_id = /$2", productID, warehouseID).Scan(&price, &discount)
		if err != nil {
			return 0, err
		}
		finalPrice := price * (1 - discount/100)
		totalPrice += finalPrice * float64(quantity)
	}
	return totalPrice, nil
}

func (r *InventoryRepository) PurchaseProducts(ctx context.Context, warehouseID uuid.UUID, products map[uuid.UUID]int) error {
	for productID, quantity := range products {
		var availableQuantity int
		err := r.db.QueryRow(ctx, "SELECT quantity FROM inventory WHERE product_id = /$1 AND warehouse_id = /$2", productID, warehouseID).Scan(&availableQuantity)
		if err != nil {
			return err
		}
		if availableQuantity < quantity {
			return fmt.Errorf("недостаточно товара: %s", productID)
		}
		_, err = r.db.Exec(ctx, "UPDATE inventory SET quantity = quantity - /$1 WHERE product_id = /$2 AND warehouse_id = /$3", quantity, productID, warehouseID)
		if err != nil {
			return err
		}
	}
	return nil
}
