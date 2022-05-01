package schema

import (
	"be_soc/graph/query"

	"github.com/graphql-go/graphql"
)

func GetAnonymousQuery(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"login": query.LoginQuery(containerRepo),
		},
	})
}

func GetClientQuery(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
<<<<<<< HEAD
			"get_novel_list":   query.GetNovelListQuery(containerRepo),
			"get_user_profile": query.GetUserProfileQuery(containerRepo),
			"get_chapter_info": query.GetChapterInfoQuery(containerRepo),
=======
			"get_novel_list":     query.GetNovelListQuery(containerRepo),
			"get_user_profile":   query.GetUserProfileQuery(containerRepo),
			"get_chapter_info":   query.GetChapterInfoQuery(containerRepo),
			"get_all_categories": query.GetAllCategories(containerRepo),
>>>>>>> 4e006f7fef8dcf1e1a07f43c8345f979cbbbbacc
		},
	})
}
func GetAdminQuery(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "query",
		Fields: graphql.Fields{},
	})
}
