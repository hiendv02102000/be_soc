package service

import "be_soc/internal/pkg/domain/domain_model/entity"

type NovelRepositoryInterface interface {
	FindNovelList(condition entity.Novels) (novels []entity.Novels, err error)
	FirstNovel(condition entity.Novels) (entity.Novels, error)
	CreateNovel(entity.Novels) (entity.Novels, error)
}
