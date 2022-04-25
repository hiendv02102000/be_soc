package mutation

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"
	"be_soc/pkg/share/middleware"
	"be_soc/pkg/share/utils"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func CreateCommentMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.CreateCommentOutput(),
		Description: "Create comment",

		Args: graphql.FieldConfigArgument{
			"comment": &graphql.ArgumentConfig{
				Type: input.CreateCommentInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["comment"].(map[string]interface{})
			createCommentReq := dto.CreateCommentRequest{
				ChapterID:      req["chapter_id"].(int),
				CommentContent: req["cmt_content"].(string),
			}
			err = utils.CheckValidate(createCommentReq)
			if err != nil {
				return
			}
			commentRepo := containerRepo["comments_repository"].(service.CommentsRepositoryInterface)
			comment, err := commentRepo.CreateComment(entity.Comments{
				ChapterID:       createCommentReq.ChapterID,
				UsersId:         user.ID,
				CommentsContent: createCommentReq.CommentContent,
			})
			if err != nil {
				return
			}
			result = map[string]interface{}{
				"id":          comment.ID,
				"cmt_content": comment.CommentsContent,
				"chapter_id":  comment.ChapterID,
				"user_id":     comment.UsersId,
			}
			return
		},
	}
}
