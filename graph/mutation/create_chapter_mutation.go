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

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func CreateChapterMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.CreateChapterOutput(),
		Description: "Create chapter",

		Args: graphql.FieldConfigArgument{
			"chapter": &graphql.ArgumentConfig{
				Type: input.CreateChapterInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			req := p.Args["chapter"].(map[string]interface{})
			createChapterReq := dto.CreateChapterRequest{
				Title:   req["title"].(string),
				NovelID: req["novel_id"].(int),
			}
			err = utils.CheckValidate(createChapterReq)
			if err != nil {
				return
			}
			user := middleware.GetUserFromContext(ctx)
			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			novel, _ := novelRepo.FirstNovel(entity.Novels{ID: createChapterReq.NovelID})
			if user.ID != novel.UsersID {
				err = errors.New("Method not allowed")
				return
			}
			chapter := entity.Chapters{
				Title:    createChapterReq.Title,
				NovelsID: createChapterReq.NovelID,
			}
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
				chapter.ContentUrl = &url
			}
			chapterRepo := containerRepo["chapters_repository"].(service.ChaptersRepositoryInterface)
			chapter, err = chapterRepo.CreateChapter(chapter)
			if err != nil {
				return
			}
			result = map[string]interface{}{
				"id":          chapter.ID,
				"title":       chapter.Title,
				"novel_id":    chapter.NovelsID,
				"content_url": chapter.ContentUrl,
				"created_at":  chapter.CreatedAt,
			}
			return
		},
	}
}
