package service

import "be_soc/internal/pkg/domain/domain_model/entity"

type NovelsCategoriesRepositoryInterface interface {
	FindNovelsCategoriesList(condition entity.NovelsCategories) (novels []entity.NovelsCategories, err error)
	CreateNovelsCategories(condition ...entity.NovelsCategories) error
}
