package service

import "be_soc/internal/pkg/domain/domain_model/entity"

type CommentsRepositoryInterface interface {
	FindComments(condition entity.Comments) (comments []entity.Comments, err error)
	FirstComment(condition entity.Comments) (entity.Comments, error)
}
