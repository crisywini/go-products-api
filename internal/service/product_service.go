package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
		return repository.Product{}, errors.New("error creating product")
	}
	return product, nil
}

func (s *ProductService) GetProduct(ctx context.Context, id uuid.UUID) (repository.Product, error) {
	product, err := s.queries.GetProduct(ctx, id)
	if err != nil {
		return repository.Product{}, errors.New("product not found")
	}
	return product, nil
}

func (s *ProductService) ListProducts(ctx context.Context) ([]repository.Product, error) {
	products, err := s.queries.ListProducts(ctx)
	if err != nil {
		return nil, errors.New("error fetching products")
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
		return repository.Product{}, errors.New("error updating product")
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	if err := s.queries.DeleteProduct(ctx, id); err != nil {
		return errors.New("error deleting product")
	}
	return nil
}
