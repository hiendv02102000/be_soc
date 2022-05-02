package query

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"

	"github.com/graphql-go/graphql"
)

func GetNovelListQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.NovelListOutput(),
		Description: "NovelListOutput",

		Args: graphql.FieldConfigArgument{
			"novel": &graphql.ArgumentConfig{
				Type: input.NovelListInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			req := p.Args["novel"].(map[string]interface{})
			NovelListReq := dto.NovelListRequest{}
			if req["id"] != nil {
				NovelListReq.ID = req["id"].(int)
			}
			if req["name"] != nil {
				NovelListReq.Name = req["name"].(string)
			}
			if req["categories"] != nil {
				NovelListReq.Categories = req["categories"].(string)
			}
			if req["user_id"] != nil {
				NovelListReq.UserID = req["user_id"].(int)
			}
			if req["is_get_chapters"] != nil {
				NovelListReq.Isgetchapters = req["is_get_chapters"].(bool)
			}
			userRepo := containerRepo["user_repository"].(service.UserRepositoryInterface)
			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			chaptersRepo := containerRepo["chapters_repository"].(service.ChaptersRepositoryInterface)
			categoriesRepo := containerRepo["categories_repository"].(service.CategoriesRepositoryInterface)
			novelscateRepo := containerRepo["novelscategories_repository"].(service.NovelsCategoriesRepositoryInterface)
			novel, err := novelRepo.FindNovelList(entity.Novels{
				ID:      NovelListReq.ID,
				Name:    NovelListReq.Name,
				UsersID: NovelListReq.UserID,
			})
			if NovelListReq.Categories != "" {
				search := []entity.Novels{}
				categories, err0 := categoriesRepo.FirstCategory(entity.Categories{
					Name: NovelListReq.Categories,
				})
				if err0 != nil {
					return
				}
				novelcate, err00 := novelscateRepo.FindNovelsCategoriesList(entity.NovelsCategories{
					CategoriesID: categories.ID,
				})

				if err00 != nil {
					return
				}
				for i := 0; i < len(novel); i++ {
					for j := 0; j < len(novelcate); j++ {
						if novel[i].ID == novelcate[j].NovelsID {
							search = append(search, novel[i])
							break
						}

					}
				}
				novel = search
			}
			novellist := make([]map[string]interface{}, 0)
			for i := 0; i < len(novel); i++ {
				cates := make([]map[string]interface{}, 0)
				chapters := make([]map[string]interface{}, 0)
				users := make([]map[string]interface{}, 0)
				nocate, err2 := novelscateRepo.FindNovelsCategoriesList(entity.NovelsCategories{
					NovelsID: novel[i].ID,
				})
				if err2 != nil {
					err = err2
					return
				}
				for i := 0; i < len(nocate); i++ {
					c, err3 := categoriesRepo.FirstCategory(entity.Categories{
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
				if NovelListReq.Isgetchapters {
					chapter, err1 := chaptersRepo.FindChaptersList(entity.Chapters{
						NovelsID: novel[i].ID,
					})
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
				user, err3 := userRepo.FirstUser(entity.Users{
					ID: novel[i].UsersID,
				})
				if err3 != nil {
					return
				}
				upro := map[string]interface{}{
					"id":         user.ID,
					"first_name": user.FirstName,
					"last_name":  user.LastName,
					"username":   user.Username,
				}
				users = append(users, upro)
				nl := map[string]interface{}{
					"id":         novel[i].ID,
					"name":       novel[i].Name,
					"img_url":    novel[i].ImageUrl,
					"user":       users,
					"chapter":    chapters,
					"categories": cates,
				}
				novellist = append(novellist, nl)
			}
			result = map[string]interface{}{
				"list": novellist,
			}
			return
		},
	}
}
