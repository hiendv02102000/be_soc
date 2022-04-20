package repository

import (
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/pkg/infrastucture/db"
)

type CategoriesRepository struct {
	DB db.Database
}

func (u *CategoriesRepository) FirstCategory(condition entity.Categories) (entity.Categories, error) {
	categories := entity.Categories{}
	err := u.DB.First(condition, &categories)
	return categories, err
}
func (u *CategoriesRepository) FindCategories(condition interface{}) (categories []entity.Categories, err error) {
	err = u.DB.Find(condition, &categories)
	return categories, err
}
func NewCategoriesRepository(db db.Database) *CategoriesRepository {
	return &CategoriesRepository{
		DB: db,
	}
}
