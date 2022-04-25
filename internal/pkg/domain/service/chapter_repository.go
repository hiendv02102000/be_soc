package service

import "be_soc/internal/pkg/domain/domain_model/entity"

type ChaptersRepositoryInterface interface {
	FindChaptersList(condition entity.Chapters) (chapters []entity.Chapters, err error)
	FirstChapter(condition entity.Chapters) (entity.Chapters, error)
	UpdateChapter(chapter, oldchapter entity.Chapters) error
	CreateChapter(chapter entity.Chapters) (entity.Chapters, error)
	DeleteChapter(chapter entity.Chapters) error
}
