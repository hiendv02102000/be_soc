package mutation

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"
	"be_soc/pkg/share/middleware"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func DeleteChapterMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.DeleteChapterOutput(),
		Description: "User Register",

		Args: graphql.FieldConfigArgument{
			"chapter": &graphql.ArgumentConfig{
				Type: input.DeleteChapterInput(),
			},
		},

		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)

			req := p.Args["chapter"].(map[string]interface{})
			delete_chapter_req := dto.DeleteChapterRequest{
				ID: req["id"].(int),
			}

			chapterRepo := containerRepo["chapter_repository"].(service.ChaptersRepositoryInterface)
			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			deleteChapter, err := chapterRepo.FirstChapter(entity.Chapters{
				ID: delete_chapter_req.ID,
			})

			novel, err := novelRepo.FirstNovel(entity.Novels{
				ID: deleteChapter.NovelsID,
			})
			if user.ID != novel.UsersID {
				err = errors.New("Method not allowed")
			}
			if err != nil {
				return
			}

			chapterRepo.DeleteChapter(entity.Chapters{
				ID: deleteChapter.ID, //dieu kien
			})

			result = map[string]interface{}{
				"id":          deleteChapter.ID,
				"title":       deleteChapter.Title,
				"content_url": deleteChapter.ContentUrl,
			}

			return
		},
	}
}
