package mutation

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"
	"be_soc/pkg/share/middleware"
	"be_soc/pkg/share/utils"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func CreateNovelMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.CreateNovelOutput(),
		Description: "Create novel",

		Args: graphql.FieldConfigArgument{
			"novel": &graphql.ArgumentConfig{
				Type: input.CreateNovelsInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			req := p.Args["novel"].(map[string]interface{})
			createNovelReq := dto.CreateNovelRequest{
				Name:         req["name"].(string),
				CategoriesID: convertCategoriesIDInputToInt(req["categories_id"].([]interface{})),
			}

			err = utils.CheckValidate(createNovelReq)
			if err != nil {
				return
			}
			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			novel, err := novelRepo.FirstNovel(entity.Novels{Name: createNovelReq.Name})
			if err != nil {
				return
			}
			if novel.ID != 0 {
				err = errors.New("Novel is already exist")
				return
			}
			user := middleware.GetUserFromContext(ctx)
			novel = entity.Novels{Name: createNovelReq.Name, UsersID: user.ID}
			file, _ := ctx.FormFile("file")
			if file != nil {
				ioFile, errFile := file.Open()
				if errFile != nil {
					err = errFile
					return
				}
				url, errUpload := utils.UploadFile(ioFile, file.Filename)
				if errUpload != nil {
					err = errUpload
					return
				}
				novel.ImageUrl = &url
			}
			categoryRepo := containerRepo["categories_repository"].(service.CategoriesRepositoryInterface)
			categoriesList, err := categoryRepo.FindCategories(createNovelReq.CategoriesID)
			if err != nil {
				return
			}

			if len(categoriesList) < len(createNovelReq.CategoriesID) {
				err = errors.New("A Category or more not exist")
				return
			}

			novel, err = novelRepo.CreateNovel(novel)
			if err != nil {
				return
			}
			novelscateRepo := containerRepo["novelscategories_repository"].(service.NovelsCategoriesRepositoryInterface)
			novelscateList := make([]entity.NovelsCategories, 0)
			resultCategories := make([]map[string]interface{}, 0)

			for _, v := range categoriesList {
				novelscateList = append(novelscateList, entity.NovelsCategories{NovelsID: novel.ID, CategoriesID: v.ID})
				resultCategories = append(resultCategories, map[string]interface{}{
					"id":   v.ID,
					"name": v.Name,
				})
			}
			err = novelscateRepo.CreateNovelsCategories(novelscateList...)

			if err != nil {
				return
			}

			result = map[string]interface{}{
				"name":       novel.Name,
				"id":         novel.ID,
				"categories": resultCategories,
				"img_url":    novel.ImageUrl,
				"created_at": novel.CreatedAt,
			}

			return
		},
	}
}
func convertCategoriesIDInputToInt(val []interface{}) []int {

	result := make([]int, len(val))
	for i, value := range val {
		switch typedValue := value.(type) {
		case int:
			result[i] = typedValue

		default:
			fmt.Println("Not an int: ", value)
		}
	}
	return result
}
