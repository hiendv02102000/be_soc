package service

import "be_soc/internal/pkg/domain/domain_model/entity"

type CategoriesRepositoryInterface interface {
	FirstCategories(condition entity.Categories) (entity.Categories, error)
	FindCategories(condition entity.Categories) (categories []entity.Categories, err error)
}
