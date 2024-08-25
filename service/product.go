package service

import (
	"context"

	"gintama/database/entity"
	"gintama/database/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, p *entity.Product) (*entity.Product, error) {
	return s.repo.CreateProduct(ctx, p)
}
