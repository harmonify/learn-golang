package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type categoryService struct {
	repository repository.ICategoryRepository
}

func (s *categoryService) Get(id string) (*entity.Category, error) {
	category := s.repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}

func NewCategoryService(
	repository repository.ICategoryRepository,
) ICategoryService {
	return &categoryService{
		repository: repository,
	}
}
