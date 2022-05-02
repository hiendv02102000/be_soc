package query

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"
	"fmt"

	"github.com/graphql-go/graphql"
)

func GetChapterInfoQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.GetChapterInfoOutput(),
		Description: "GetChapterInfoOutput",

		Args: graphql.FieldConfigArgument{
			"chapter": &graphql.ArgumentConfig{
				Type: input.GetChapterInfoInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			req := p.Args["chapter"].(map[string]interface{})
			getchapterinfo := dto.GetChapterInfoRequest{
				ID: req["id"].(int),
			}
			if err != nil {
				return
			}
			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			gciRepo := containerRepo["chapter_repository"].(service.ChaptersRepositoryInterface)
			commentRepo := containerRepo["comments_repository"].(service.CommentsRepositoryInterface)
			chapterinfo, err := gciRepo.FirstChapter(entity.Chapters{
				ID: getchapterinfo.ID,
			})
			if err != nil {
				return
			}
			cmtchapter, err := commentRepo.FindComments(entity.Comments{
				ChapterID: chapterinfo.ID,
			})
			comments := make([]map[string]interface{}, 0)
			for i := 0; i < len(cmtchapter); i++ {
				c, err1 := commentRepo.FirstComment(entity.Comments{
					ID: cmtchapter[i].ID,
				})
				if err1 != nil {
					return
				}
				comment := map[string]interface{}{
					"id":              c.ID,
					"user_id":         c.UsersId,
					"comment_content": c.CommentsContent,
				}
				comments = append(comments, comment)
			}
			result = map[string]interface{}{
				"id":          chapterinfo.ID,
				"novel_id":    chapterinfo.NovelsID,
				"title":       chapterinfo.Title,
				"content_url": chapterinfo.ContentUrl,
				"comments":    comments,
			}
			novelRepo.UpdateView(
				"update novels set view = view +1 where id =" + fmt.Sprint(chapterinfo.NovelsID) + ";",
			)
			return
		},
	}
}
