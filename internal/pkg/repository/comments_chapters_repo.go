package repository

import (
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/pkg/infrastucture/db"
)

type CommentsChaptersRepository struct {
	DB db.Database
}

func (u *CommentsChaptersRepository) FindCommentsChapters(condition entity.CommentsChapters) (commentschapters []entity.CommentsChapters, err error) {
	err = u.DB.Find(condition, &commentschapters)
	return
}
func NewCommentsChaptersRepository(db db.Database) *CommentsChaptersRepository {
	return &CommentsChaptersRepository{
		DB: db,
	}
}
