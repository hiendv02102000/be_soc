package query

import (
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"

	"github.com/graphql-go/graphql"
)

func GetAllCategories(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.GetAllCategoriesOutput(),
		Description: "Categories",

		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			// ctx := p.Context.(*gin.Context)
			// user := middleware.GetUserFromContext(ctx)

			if err != nil {
				return
			}
			cateRepo := containerRepo["categories_repository"].(service.CategoriesRepositoryInterface)

			allCate, err := cateRepo.FindCategories(entity.Categories{})
			categories := make([]map[string]interface{}, 0)

			for i := 0; i < len(allCate); i++ {

				cate := map[string]interface{}{
					"id":   allCate[i].ID,
					"name": allCate[i].Name,
				}
				categories = append(categories, cate)
			}
			result = map[string]interface{}{
				"list": categories,
			}
			return
		},
	}
}
