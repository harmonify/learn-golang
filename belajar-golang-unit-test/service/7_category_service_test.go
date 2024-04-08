package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockCategoryRepository = repository.NewMockCategoryRepository()
var serviceUnderTest = NewCategoryService(mockCategoryRepository)

func TestCategoryService_GetNotFound(t *testing.T) {
	// program mock
	mockCategoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := serviceUnderTest.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	mockCategoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := serviceUnderTest.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
