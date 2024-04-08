package repository

import (
	"belajar-golang-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

type mockCategoryRepository struct {
	Mock mock.Mock
}

func (repository *mockCategoryRepository) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}

func NewMockCategoryRepository() *mockCategoryRepository {
	// return &categoryRepositoryMock{Mock: mock.Mock{}} // is this the same as the below one?
	return &mockCategoryRepository{}
}
