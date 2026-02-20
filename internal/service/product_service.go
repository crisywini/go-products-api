package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/crisywini/go-products-api/internal/repository"
	"github.com/google/uuid"
)

type ProductService struct {
	queries *repository.Queries
}

func NewProductService(queries *repository.Queries) *ProductService {
	return &ProductService{queries: queries}
}

func (s *ProductService) CreateProduct(ctx context.Context, name, description string, price float64, stock int32, userID uuid.UUID) (repository.Product, error) {
	product, err := s.queries.CreateProduct(ctx, repository.CreateProductParams{
		Name:        name,
		Description: sql.NullString{String: description, Valid: description != ""},
		Price:       fmt.Sprintf("%.2f", price),
		Stock:       stock,
		UserID:      userID,
	})
	if err != nil {
		log.Printf("[ProductService.CreateProduct] db error: %v", err)
		return repository.Product{}, fmt.Errorf("error creating product: %w", err)
	}
	return product, nil
}

func (s *ProductService) GetProduct(ctx context.Context, id uuid.UUID) (repository.Product, error) {
	product, err := s.queries.GetProduct(ctx, id)
	if err != nil {
		log.Printf("[ProductService.GetProduct] db error: %v", err)
		return repository.Product{}, fmt.Errorf("product not found: %w", err)
	}
	return product, nil
}

func (s *ProductService) ListProducts(ctx context.Context) ([]repository.Product, error) {
	products, err := s.queries.ListProducts(ctx)
	if err != nil {
		log.Printf("[ProductService.ListProducts] db error: %v", err)
		return nil, fmt.Errorf("error fetching products: %w", err)
	}
	if products == nil {
		products = []repository.Product{}
	}
	return products, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id uuid.UUID, name, description string, price float64, stock int32) (repository.Product, error) {
	product, err := s.queries.UpdateProduct(ctx, repository.UpdateProductParams{
		ID:          id,
		Name:        name,
		Description: sql.NullString{String: description, Valid: description != ""},
		Price:       fmt.Sprintf("%.2f", price),
		Stock:       stock,
	})
	if err != nil {
		log.Printf("[ProductService.UpdateProduct] db error: %v", err)
		return repository.Product{}, fmt.Errorf("error updating product: %w", err)
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	if err := s.queries.DeleteProduct(ctx, id); err != nil {
		log.Printf("[ProductService.DeleteProduct] db error: %v", err)
		return fmt.Errorf("error deleting product: %w", err)
	}
	return nil
}
