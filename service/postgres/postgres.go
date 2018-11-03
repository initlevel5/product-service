package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/initlevel5/product-service/service"
	_ "github.com/lib/pq"
)

var (
	errNotImplemented = errors.New("postgres: not implemented")
)

// productService represents a PostgreSQL implementation of service.ProductService interface
type productService struct {
	db *sql.DB
	*log.Logger
}

func NewProductService(logger *log.Logger) *productService {
	return &productService{Logger: logger}
}

func (s *productService) CreateProduct(ctx context.Context, catid, title string, price float64, manufacturer, description string) (string, error) {
	return "", errNotImplemented
}

func (s *productService) UpdateProduct(ctx context.Context, id string, price float64) (string, error) {
	return "", errNotImplemented
}

func (s *productService) DeleteProduct(ctx context.Context, id string) (string, error) {
	return "", errNotImplemented
}

func (s *productService) Product(ctx context.Context, id string) (*service.Product, error) {
	return nil, errNotImplemented
}

func (s *productService) SearchProduct(ctx context.Context, title string) (string, error) {
	return "", errNotImplemented
}
