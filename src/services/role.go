package services

import (
	"gintama/src/models"
	"gintama/src/repository"
	"gintama/src/schema"
)

type IService interface {
	GetAll(queryFiltered models.QueryFiltered) ([]models.Role, int64, error)
	CreateRole(input schema.RoleSchema) (models.Role, error)
}

type service struct {
	repository repository.IRepository
}

func RoleService(repository repository.IRepository) *service {
	return &service{repository}
}

// Get All
func (s *service) GetAll(queryFiltered models.QueryFiltered) ([]models.Role, int64, error) {
	data, total, err := s.repository.GetAll(queryFiltered)
	if err != nil {
		return data, total, err
	}

	return data, total, nil
}

// Create
func (s *service) CreateRole(input schema.RoleSchema) (models.Role, error) {
	formData := models.Role{}
	formData.Name = input.Name

	data, err := s.repository.Create(formData)
	if err != nil {
		return data, err
	}

	return data, nil
}
