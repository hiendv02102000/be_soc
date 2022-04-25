package repository

import (
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/pkg/infrastucture/db"
)

type NovelRepository struct {
	DB db.Database
}

func (u *NovelRepository) FirstNovel(condition entity.Novels) (entity.Novels, error) {
	novel := entity.Novels{}
	err := u.DB.First(condition, &novel)
	return novel, err
}
func (u *NovelRepository) FindNovelList(condition entity.Novels) (novel []entity.Novels, err error) {
	err = u.DB.Find(condition, &novel)
	return novel, err
}
func (u *NovelRepository) CreateNovel(novel entity.Novels) (entity.Novels, error) {
	err := u.DB.Create(&novel)
	return novel, err
}
func (u *NovelRepository) UpdateNovel(novel, oldNovel entity.Novels) error {
	return u.DB.Update(entity.Novels{}, oldNovel, novel)
}
func NewNovelRepository(db db.Database) *NovelRepository {
	return &NovelRepository{
		DB: db,
	}
}
func (u *NovelRepository) DeleteNovel(novel entity.Novels) error {
	err := u.DB.Delete(&novel)
	return err
}
