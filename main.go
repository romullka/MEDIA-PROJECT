package main

import (
	"context"
	"log"
	"net/http"

	"github.com/romullka/MEDIA-PROJECT/internal/handler"
	"github.com/romullka/MEDIA-PROJECT/internal/repository"
	"github.com/romullka/MEDIA-PROJECT/internal/service"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func main() {
	db, err := pgx.Connect(context.Background(), "postgres://user:password@postgres:5432/warehouse_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	r := mux.NewRouter()

	warehouseRepo := repository.NewWarehouseRepository(db)
	productRepo := repository.NewProductRepository(db)
	inventoryRepo := repository.NewInventoryRepository(db)
	analyticsRepo := repository.NewAnalyticsRepository(db)

	inventoryService := service.NewInventoryService(inventoryRepo, analyticsRepo)
	analyticsService := service.NewAnalyticsService(analyticsRepo)

	warehouseHandler := handler.NewWarehouseHandler(warehouseRepo)
	productHandler := handler.NewProductHandler(productRepo)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsService)

	r.HandleFunc("/api/health", healthCheckHandler).Methods("GET")
	r.HandleFunc("/api/warehouses", warehouseHandler.CreateWarehouse).Methods("POST")
	r.HandleFunc("/api/warehouses", warehouseHandler.GetWarehouses).Methods("GET")
	r.HandleFunc("/api/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products", productHandler.GetProducts).Methods("GET")

	r.HandleFunc("/api/inventory", inventoryHandler.CreateInventory).Methods("POST")
	r.HandleFunc("/api/purchase", inventoryHandler.PurchaseProducts).Methods("POST")

	r.HandleFunc("/api/analytics", analyticsHandler.GetAnalyticsByWarehouse).Methods("GET")
	r.HandleFunc("/api/top-warehouses", analyticsHandler.GetTopWarehouses).Methods("GET")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
