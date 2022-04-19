package mutation

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"

	"github.com/graphql-go/graphql"
)

func ListNovelsMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.ListNovelsOutput(),
		Description: "ListNovelOutput",

		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.ListNovelsInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			//arr := make([]map[string]interface{}, 0)
			req := p.Args["user"].(map[string]interface{})
			listNovelsReq := dto.ListNovelsRequest{
				Name:          req["name"].(string),
				Categories:    req["categories"].(string),
				UserID:        req["user_id"].(int),
				Isgetchapters: req["is_get_chapters"].(bool),
			}
			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			chaptersRepo := containerRepo["chapters_repository"].(service.ChaptersRepositoryInterface)
			categoriesRepo := containerRepo["categories_repository"].(service.CategoriesRepositoryInterface)
			novelscateRepo := containerRepo["novelscategories_repository"].(service.NovelsCategoriesRepositoryInterface)
			categories, err := categoriesRepo.FirstCategories(entity.Categories{
				Name: listNovelsReq.Categories,
			})
			novelcate, err := novelscateRepo.FindNovelsCategoriesList(entity.NovelsCategories{
				CategoriesID: categories.ID,
			})
			novel, err := novelRepo.FindNovelList(entity.Novels{
				Name:    listNovelsReq.Name,
				UsersID: listNovelsReq.UserID,
				ID:      novelcate.ID,
			})

			listnovels := make([]map[string]interface{}, 0)
			for i := 0; i < len(novel); i++ {
				cates := make([]map[string]interface{}, 0)
				chapters := make([]map[string]interface{}, 0)
				for i := 0; i < len(novel); i++ {
					nocate, err2 := novelscateRepo.FindNovelsCategoriesList(entity.NovelsCategories{
						NovelsID: novel[i].ID,
					})
					if err2 != nil {
						err = err2
						return
					}
					for i := 0; i < len(nocate); i++ {
						c, err3 := categoriesRepo.FirstCategories(entity.Categories{
							ID: nocate[i].CategoriesID,
						})
						if err2 != nil {
							err = err3
							return
						}
						categories := map[string]interface{}{
							"name": c.Name,
						}
						cates = append(cates, categories)
					}
				}
				if listNovelsReq.Isgetchapters {
					chapter, err1 := chaptersRepo.FindChaptersList(entity.Chapters{
						NovelsID: novel[i].ID,
					})
					//fmt.Println(chapter)
					if err1 != nil {
						err = err1
						return
					}
					for i := 0; i < len(chapter); i++ {
						chapterco := map[string]interface{}{
							"id":          chapter[i].ID,
							"title":       chapter[i].Title,
							"content_url": chapter[i].ContentUrl,
						}
						chapters = append(chapters, chapterco)
					}
				}
				listnovel := map[string]interface{}{
					"id":         novel[i].ID,
					"name":       novel[i].Name,
					"user_id":    novel[i].UsersID,
					"chapter":    chapters,
					"categories": cates,
				}
				listnovels = append(listnovels, listnovel)
			}
			result = map[string]interface{}{
				"list": listnovels,
			}
			return
		},
	}
}
