package schema

import (
	"be_soc/graph/mutation"

	"github.com/graphql-go/graphql"
)

func GetAnonymousMutation(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"register_user": mutation.CreateUserMutation(containerRepo),
		},
	})
}

func GetClientMutation(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"change_password": mutation.ChangePasswordMutation(containerRepo),
			"change_profile":  mutation.ChangeProfileMuitation(containerRepo),
			"create_novel":    mutation.CreateNovelMutation(containerRepo),
			"update_novel":    mutation.UpdateNovelMutation(containerRepo),
			"update_chapter":  mutation.UpdateChapterMutation(containerRepo),
			"create_chapter":  mutation.CreateChapterMutation(containerRepo),
			"delete_novel":    mutation.DeleteNovelMutation(containerRepo),
			"delete_chapter":  mutation.DeleteChapterMutation(containerRepo),
		},
	})
}
func GetAdminMutation(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "query",
		Fields: graphql.Fields{},
	})
}
