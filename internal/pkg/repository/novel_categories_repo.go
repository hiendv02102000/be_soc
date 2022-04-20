package repository

import (
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/pkg/infrastucture/db"
)

type NovelsCategoriesRepository struct {
	DB db.Database
}

func (u *NovelsCategoriesRepository) FindNovelsCategoriesList(condition entity.NovelsCategories) (novelscategories []entity.NovelsCategories, err error) {
	err = u.DB.Find(condition, &novelscategories)
	return
}
func (u *NovelsCategoriesRepository) CreateNovelsCategories(value ...entity.NovelsCategories) (err error) {
	for _, v := range value {
		err = u.DB.Create(&v)
		if err != nil {
			break
		}
	}

	return err
}
func NewNovelsCategoriesRepository(db db.Database) *NovelsCategoriesRepository {
	return &NovelsCategoriesRepository{
		DB: db,
	}
}
