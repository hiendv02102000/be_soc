package schema

import (
	"be_soc/internal/pkg/repository"
	"be_soc/pkg/infrastucture/db"
)

func GetContainerRepo(data db.Database) map[string]interface{} {

	return map[string]interface{}{
		"user_repository":             repository.NewUserRepository(data),
		"novel_repository":            repository.NewNovelRepository(data),
		"chapters_repository":         repository.NewChaptersRepository(data),
		"novelscategories_repository": repository.NewNovelsCategoriesRepository(data),
		"categories_repository":       repository.NewCategoriesRepository(data),
		"chapter_repository":          repository.NewChaptersRepository(data),
	}
}
