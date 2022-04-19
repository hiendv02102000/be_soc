package mutation

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/pkg/share/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func CreateNovelMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.CreateUserOutput(),
		Description: "Create novel",

		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.CreateNovelsInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			file, _ := ctx.FormFile("file")
			novel := entity.Novels{}
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

			fmt.Println(novel)
			return
		},
	}
}
