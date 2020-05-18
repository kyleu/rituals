package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	profileResolver       Callback
	profileSprintResolver Callback
	profileType           *graphql.Object
)

func initProfile() {
	profileResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.Profile.ToProfile(), nil
	}

	profileSprintResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.GetByMember(ctx.Profile.UserID, nil)
	}

	profileType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Profile",
			Fields: graphql.Fields{
				"userID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"role": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"theme": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"navColor": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"linkColor": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"picture": &graphql.Field{
					Type: graphql.String,
				},
				"locale": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"sprints": &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(sprintType)),
					Description: "Your current sprints",
					Resolve:     ctxF(profileSprintResolver),
				},
			},
		},
	)
}
