package repository

import (
	"context"
	"fmt"

	"gintama/database/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) CreateProduct(ctx context.Context, p *entity.Product) (*entity.Product, error) {
	// assign value
	p.ID = uuid.New()

	err := repo.db.Create(&p).Error
	if err != nil {
		return nil, fmt.Errorf("error inserting product: %w", err)
	}

	return p, nil
}
