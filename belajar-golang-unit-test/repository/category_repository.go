package repository

import "belajar-golang-unit-test/entity"

type categoryRepository struct{}

func (r *categoryRepository) FindById(id string) *entity.Category {
	return &entity.Category{} // only serves as an example
}

func NewCategoryRepository() ICategoryRepository {
	return &categoryRepository{}
}
