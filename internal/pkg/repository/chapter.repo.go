package repository

import (
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/pkg/infrastucture/db"
)

type ChaptersRepository struct {
	DB db.Database
}

func (u *ChaptersRepository) FindChaptersList(condition entity.Chapters) (chapters []entity.Chapters, err error) {
	err = u.DB.Find(condition, &chapters)
	return
}
func NewChaptersRepository(db db.Database) *ChaptersRepository {
	return &ChaptersRepository{
		DB: db,
	}
}
