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

func UpdateChapterMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.UpdateChapterOutput(),
		Description: "Update chapter",

		Args: graphql.FieldConfigArgument{
			"chapter": &graphql.ArgumentConfig{
				Type: input.UpdateChapterInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["chapter"].(map[string]interface{})
			updateChapterReq := dto.UpdateChapterRequest{
				NovelID:   req["novel_id"].(int),
				ChapterID: req["chapter_id"].(int),
			}
			if req["title"] != nil {
				updateChapterReq.Title = req["title"].(string)
			}

			chapterRepo := containerRepo["chapter_repository"].(service.ChaptersRepositoryInterface)
			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			novel, err := novelRepo.FirstNovel(entity.Novels{
				ID: updateChapterReq.NovelID,
			})
			if err != nil {
				return
			}
			if user.ID != novel.UsersID {
				err = errors.New("Method not allowed")
			}
			chapter, err := chapterRepo.FirstChapter(entity.Chapters{
				ID: updateChapterReq.ChapterID,
			})
			if chapter.NovelsID != updateChapterReq.NovelID {
				err = errors.New("Method not allowed")
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
				updateChapterReq.Contenturl = &url
			}
			chapterRepo.UpdateChapter(entity.Chapters{
				Title:      updateChapterReq.Title,
				ContentUrl: updateChapterReq.Contenturl,
			}, chapter)
			newchapter, err := chapterRepo.FirstChapter(entity.Chapters{
				ID: updateChapterReq.ChapterID,
			})
			result = map[string]interface{}{
				"id":          newchapter.ID,
				"title":       newchapter.Title,
				"content_url": newchapter.ContentUrl,
				"update_at":   newchapter.UpdatedAt,
			}
			// timeNow := time.Now()

			// newUser := entity.Users{Username: loginReq.Username,
			// 	Password: loginReq.Password}
			return
		},
	}
}
