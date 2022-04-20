package service

import "be_soc/internal/pkg/domain/domain_model/entity"

type CategoriesRepositoryInterface interface {
	FirstCategory(condition entity.Categories) (entity.Categories, error)
	FindCategories(condition interface{}) (categories []entity.Categories, err error)
}
