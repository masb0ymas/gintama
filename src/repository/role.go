package repository

import (
	"gintama/src/models"
	"strconv"

	"gorm.io/gorm"
)

type IRepository interface {
	GetAll(queryFiltered models.QueryFiltered) ([]models.Role, int64, error)
	FindById(id int) (models.Role, error)
	Create(role models.Role) (models.Role, error)
}

type repository struct {
	db *gorm.DB
}

func RoleRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Get All
func (r *repository) GetAll(queryFiltered models.QueryFiltered) ([]models.Role, int64, error) {
	var data []models.Role
	var count int64

	queryPage, _ := strconv.Atoi(queryFiltered.Page)
	queryPageSize, _ := strconv.Atoi(queryFiltered.PageSize)

	page := queryPage | 1
	pageSize := queryPageSize | 10

	err := r.db.Model(models.Role{}).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&data).Error

	// total
	r.db.Model(models.Role{}).Count(&count)

	if err != nil {
		return data, count, err
	}

	return data, count, nil
}

// Find By Id
func (r *repository) FindById(id int) (models.Role, error) {
	var data models.Role

	err := r.db.Model(models.Role{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

// Create
func (r *repository) Create(data models.Role) (models.Role, error) {
	err := r.db.Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}
