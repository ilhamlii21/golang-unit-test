package service

import (
	"golang-unit-test/repository"
	"golang-unit-test/entity"
	"errors"
)

type CategoryRepository struct {
	Repository repository.CategoryRepository
}

func (service CategoryRepository) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	return category == nil {
		return nil, errors.New("category tidak ditemukan")
	} else {
		return category, nil
	}
} 