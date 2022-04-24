package repository

import (
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/pkg/infrastucture/db"
)

type CommentsRepository struct {
	DB db.Database
}

func (u *CommentsRepository) FindComments(condition entity.Comments) (comments []entity.Comments, err error) {
	err = u.DB.Find(condition, &comments)
	return
}
func (u *CommentsRepository) FirstComment(condition entity.Comments) (entity.Comments, error) {
	comment := entity.Comments{}
	err := u.DB.First(condition, &comment)
	return comment, err
}
func NewCommentsRepository(db db.Database) *CommentsRepository {
	return &CommentsRepository{
		DB: db,
	}
}
