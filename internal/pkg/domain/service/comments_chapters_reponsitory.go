package service

import "be_soc/internal/pkg/domain/domain_model/entity"

type CommentsChaptersRepositoryInterface interface {
	FindCommentsChapters(condition entity.CommentsChapters) (commentschapters []entity.CommentsChapters, err error)
}
