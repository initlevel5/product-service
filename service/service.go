package service

import (
	"context"
	"time"
)

type Product struct {
	ID           string
	CatID        string
	Title        string
	Price        float64
	Manufacturer string
	Description  string
	Created      time.Time
	Updated      time.Time
}

type ProductService interface {
	CreateProduct(ctx context.Context, catid, title string, price float64, manufacturer, description string) (string, error)
	UpdateProduct(ctx context.Context, id string, price float64) (string, error)
	DeleteProduct(ctx context.Context, id string) (string, error)
	Product(ctx context.Context, id string) (*Product, error)
	SearchProduct(ctx context.Context, title string) (string, error)
}
