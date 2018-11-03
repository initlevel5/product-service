package graphql

import (
	"context"
	"log"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/initlevel5/product-service/service"
)

type resolver struct {
	s service.ProductService
	*log.Logger
}

func NewResolver(s service.ProductService, logger *log.Logger) *resolver {
	return &resolver{s: s, Logger: logger}
}

func (r *resolver) Product(args struct{ ID graphql.ID }) *productResolver {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	p, err := r.s.Product(ctx, string(args.ID))
	if err != nil {
		r.Printf("Product(): %v", err)
		return nil
	}
	return &productResolver{p}
}

func (r *resolver) SearchProduct(args struct{ Title string }) *graphql.ID {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	id, err := r.s.SearchProduct(ctx, args.Title)
	if err != nil {
		r.Printf("SearchProduct(): %v", err)
		return nil
	}

	gid := new(graphql.ID)
	*gid = graphql.ID(id)
	return gid
}

func (r *resolver) CreateProduct(args struct {
	CatID        string
	Title        string
	Price        float64
	Manufacturer string
	Description  *string
}) *graphql.ID {

	var descr string

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if args.Description != nil {
		descr = *args.Description
	}

	id, err := r.s.CreateProduct(ctx, args.CatID, args.Title, args.Price, args.Manufacturer, descr)
	if err != nil {
		r.Printf("CreateProduct(): %v", err)
		return nil
	}

	gid := new(graphql.ID)
	*gid = graphql.ID(id)
	return gid
}

func (r *resolver) UpdateProduct(args struct {
	ID    graphql.ID
	Price float64
}) *graphql.ID {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	id, err := r.s.UpdateProduct(ctx, string(args.ID), args.Price)
	if err != nil {
		r.Printf("UpdateProduct(): %v", err)
		return nil
	}

	gid := new(graphql.ID)
	*gid = graphql.ID(id)
	return gid
}

func (r *resolver) DeleteProduct(args struct{ ID graphql.ID }) *graphql.ID {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	id, err := r.s.DeleteProduct(ctx, string(args.ID))
	if err != nil {
		r.Printf("DeleteProduct(): %v", err)
		return nil
	}

	gid := new(graphql.ID)
	*gid = graphql.ID(id)
	return gid
}

type productResolver struct {
	p *service.Product
}

func (r *productResolver) ID() graphql.ID {
	return graphql.ID(r.p.ID)
}

func (r *productResolver) CatID() graphql.ID {
	return graphql.ID(r.p.CatID)
}

func (r *productResolver) Title() string {
	return r.p.Title
}

func (r *productResolver) Price() float64 {
	return r.p.Price
}

func (r *productResolver) Manufacturer() string {
	return r.p.Manufacturer
}

func (r *productResolver) Description() *string {
	if r.p.Description == "" {
		return nil
	}
	return &r.p.Description
}

func (r *productResolver) Created() string {
	return r.p.Created.String()
}

func (r *productResolver) Updated() string {
	return r.p.Updated.String()
}
