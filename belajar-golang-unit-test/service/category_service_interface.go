package service

import (
	"belajar-golang-unit-test/entity"
)

type ICategoryService interface {
	Get(id string) (*entity.Category, error)
}
