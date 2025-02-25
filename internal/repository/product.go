package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/romullka/MEDIA-PROJECT/internal/model"
)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, name, description string, specs map[string]string, weight float64, barcode string) (model.Product, error) {
	id := uuid.New()
	_, err := r.db.Exec(ctx, "INSERT INTO products (id, name, description, specs, weight, barcode) VALUES (/$1, /$2, /$3, /$4, /$5, /$6)", id, name, description, specs, weight, barcode)
	if err != nil {
		return model.Product{}, err
	}
	return model.Product{ID: id, Name: name, Description: description, Specs: specs, Weight: weight, Barcode: barcode}, nil
}

func (r *ProductRepository) GetProducts(ctx context.Context) ([]model.Product, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, description, specs, weight, barcode FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Specs, &product.Weight, &product.Barcode); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id uuid.UUID) (model.Product, error) {
	var product model.Product
	err := r.db.QueryRow(ctx, "SELECT id, name, description, specs, weight, barcode FROM products WHERE id = /$1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Specs, &product.Weight, &product.Barcode)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id uuid.UUID, name, description string, specs map[string]string, weight float64, barcode string) (model.Product, error) {
	_, err := r.db.Exec(ctx, "UPDATE products SET name = /$1, description =/$2, specs = /$3, weight = /$4, barcode = /$5 WHERE id = /$6", name, description, specs, weight, barcode, id)
	if err != nil {
		return model.Product{}, err
	}
	return model.Product{ID: id, Name: name, Description: description, Specs: specs, Weight: weight, Barcode: barcode}, nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, "DELETE FROM products WHERE id = /$1", id)
	return err
}
