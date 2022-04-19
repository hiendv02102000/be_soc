package service

import "be_soc/internal/pkg/domain/domain_model/entity"

type ChaptersRepositoryInterface interface {
	FindChaptersList(condition entity.Chapters) (chapters []entity.Chapters, err error)
}
