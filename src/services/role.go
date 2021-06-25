package services

import (
	"gintama/src/models"
	"gintama/src/repository"
	"gintama/src/schema"
)

type IService interface {
	GetAll(queryFiltered models.QueryFiltered) ([]models.Role, int64, error)
	FindById(input schema.RoleByIdSchema) (models.Role, error)
	CreateRole(input schema.RoleSchema) (models.Role, error)
	Update(params schema.RoleByIdSchema, input schema.RoleSchema) (models.Role, error)
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

// Find By Id
func (s *service) FindById(input schema.RoleByIdSchema) (models.Role, error) {
	data, err := s.repository.FindById(input.ID)
	if err != nil {
		return data, err
	}

	return data, nil
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

// Update
func (s *service) Update(params schema.RoleByIdSchema, input schema.RoleSchema) (models.Role, error) {
	data, err := s.repository.FindById(params.ID)
	if err != nil {
		return data, err
	}

	data.Name = input.Name

	updateData, err := s.repository.Update(data)
	if err != nil {
		return updateData, err
	}

	return updateData, nil
}
