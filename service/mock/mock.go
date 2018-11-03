package mock

import (
	"context"
	"errors"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/initlevel5/product-service/service"
)

var (
	errAlreadyExists = errors.New("mock: already exists")
	errNotFound      = errors.New("mock: not found")
)

var (
	defaultProducts = []*service.Product{
		{
			ID:           "1001",
			CatID:        "101",
			Title:        "Socks",
			Price:        2.95,
			Manufacturer: "Adidas",
			Description:  "Best socks",
		},
		{
			ID:           "1002",
			CatID:        "101",
			Title:        "Jeans",
			Price:        20.99,
			Manufacturer: "Levi's",
			Description:  "Good jeans",
		},
		{
			ID:           "1003",
			CatID:        "101",
			Title:        "T-shirt",
			Price:        7.45,
			Manufacturer: "Ostin",
		},
	}
)

// productService represents a mock implementation of service.ProductService interface
type productService struct {
	mu *sync.RWMutex
	db map[string]*service.Product
	*log.Logger
}

func NewProductService(logger *log.Logger) *productService {
	s := &productService{
		mu:     &sync.RWMutex{},
		db:     make(map[string]*service.Product),
		Logger: logger,
	}

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)

	for _, p := range defaultProducts {
		p.Created = now
		s.db[p.ID] = p
	}
	return s
}

func (s *productService) CreateProduct(ctx context.Context, catid, title string, price float64, manufacturer, description string) (string, error) {
	_, err := s.SearchProduct(ctx, title)
	if err == nil {
		err = errAlreadyExists
	}
	if err != nil && err != errNotFound {
		return "", err
	}

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)

	p := &service.Product{
		ID:           strconv.FormatInt(int64(len(s.db)+1), 10),
		CatID:        catid,
		Title:        title,
		Price:        price,
		Manufacturer: manufacturer,
		Description:  description,
		Created:      now,
		Updated:      now,
	}

	s.mu.Lock()
	s.db[p.ID] = p
	s.mu.Unlock()

	return p.ID, nil
}

func (s *productService) UpdateProduct(ctx context.Context, id string, price float64) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if p, ok := s.db[id]; ok {
		p.Price = price
		loc, _ := time.LoadLocation("UTC")
		p.Updated = time.Now().In(loc)
		return id, nil
	}
	return "", errNotFound
}

func (s *productService) DeleteProduct(ctx context.Context, id string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.db[id]; ok {
		delete(s.db, id)
		return id, nil
	}
	return "", errNotFound
}

func (s *productService) Product(ctx context.Context, id string) (*service.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if p, ok := s.db[id]; ok {
		return &service.Product{
			ID:           p.ID,
			Title:        p.Title,
			Price:        p.Price,
			Manufacturer: p.Manufacturer,
			Description:  p.Description,
			Created:      p.Created,
			Updated:      p.Updated,
		}, nil
	}
	return nil, errNotFound
}

func (s *productService) SearchProduct(ctx context.Context, title string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, p := range s.db {
		if p.Title == title {
			return p.ID, nil
		}
	}
	return "", errNotFound
}
