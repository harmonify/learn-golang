package repository

import "belajar-golang-unit-test/entity"

type ICategoryRepository interface {
	FindById(id string) *entity.Category
}
