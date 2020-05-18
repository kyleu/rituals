package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	retroArgs      graphql.FieldConfigArgument
	retroResolver  Callback
	retrosResolver Callback
	retroType      *graphql.Object
)

func initRetro() {
	retroArgs = graphql.FieldConfigArgument{
		"key": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	retroResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		slug, ok := p.Args["key"].(string)
		if ok {
			return ctx.App.Retro.GetBySlug(slug)
		}
		return nil, nil
	}

	retrosResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.List(paramSetFromGraphQLParams("retro", params))
	}

	retroType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Retro",
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
				"sprintID": &graphql.Field{
					Type: graphql.String,
				},
				"owner": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"status": &graphql.Field{
					Type: graphql.String,
				},
				"categories": &graphql.Field{
					Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
				},
				"created": &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)
}
