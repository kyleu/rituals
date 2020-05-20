package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	estimateArgs           graphql.FieldConfigArgument
	estimateResolver       Callback
	estimatesResolver      Callback
	estimateMemberResolver Callback
	estimateTeamResolver   Callback
	estimateSprintResolver Callback
	estimateType           *graphql.Object
)

func initEstimate() {
	estimateArgs = graphql.FieldConfigArgument{
		"key": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	estimateStatusType := graphql.NewEnum(graphql.EnumConfig{
		Name: "EstimateStatus",
		Values: graphql.EnumValueConfigMap{
			"new":      &graphql.EnumValueConfig{Value: "new"},
			"active":   &graphql.EnumValueConfig{Value: "active"},
			"complete": &graphql.EnumValueConfig{Value: "complete"},
			"deleted":  &graphql.EnumValueConfig{Value: "deleted"},
		},
	})

	estimateResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		slug, ok := p.Args["key"].(string)
		if ok {
			return ctx.App.Estimate.GetBySlug(slug)
		}
		return nil, nil
	}

	estimatesResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Estimate.List(paramSetFromGraphQLParams(util.SvcEstimate.Key, params, ctx.Logger))
	}

	estimateMemberResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Estimate.Members.GetByModelID(p.Source.(*estimate.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger)), nil
	}

	estimateTeamResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		sess := p.Source.(*estimate.Session)
		if sess.TeamID != nil {
			return ctx.App.Team.GetByID(*sess.TeamID)
		}
		return nil, nil
	}

	estimateSprintResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		sess := p.Source.(*estimate.Session)
		if sess.SprintID != nil {
			return ctx.App.Sprint.GetByID(*sess.SprintID)
		}
		return nil, nil
	}

	estimateType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Estimate",
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
				"teamID": &graphql.Field{
					Type: graphql.String,
				},
				"sprintID": &graphql.Field{
					Type: graphql.String,
				},
				"owner": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"status": &graphql.Field{
					Type: graphql.NewNonNull(estimateStatusType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*estimate.Session).Status.Key, nil
					},
				},
				"choices": &graphql.Field{
					Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
				},
				"created": &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"members": &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This estimate's members",
					Args:        listArgs,
					Resolve:     ctxF(estimateMemberResolver),
				},
			},
		},
	)
}
