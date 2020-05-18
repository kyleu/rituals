package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	sprintArgs      graphql.FieldConfigArgument
	sprintResolver  Callback
	sprintsResolver Callback
	sprintType      *graphql.Object
)

func initSprint() {
	sprintArgs = graphql.FieldConfigArgument{
		"key": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	sprintResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		slug, ok := p.Args["key"].(string)
		if ok {
			return ctx.App.Sprint.GetBySlug(slug)
		}
		return nil, nil
	}

	sprintsResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.List(paramSetFromGraphQLParams("sprint", params))
	}

	sprintType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Sprint",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"slug": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"title": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"owner": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"endDate": &graphql.Field{
					Type: graphql.String,
				},
				"created": &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)
}
