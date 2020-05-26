package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	estimateArgs               graphql.FieldConfigArgument
	estimateResolver           Callback
	estimatesResolver          Callback
	estimateMemberResolver     Callback
	estimatePermissionResolver Callback
	estimateTeamResolver       Callback
	estimateSprintResolver     Callback
	estimateType               *graphql.Object
)

func initEstimate() {
	estimateArgs = graphql.FieldConfigArgument{
		util.KeyKey: &graphql.ArgumentConfig{
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
		slug, err := paramKeyString(p)
		if err != nil {
			return nil, err
		}
		return ctx.App.Estimate.GetBySlug(slug)
	}

	estimatesResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Estimate.List(paramSetFromGraphQLParams(util.SvcEstimate.Key, params, ctx.Logger))
	}

	estimateMemberResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Estimate.Members.GetByModelID(p.Source.(*estimate.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger)), nil
	}

	estimatePermissionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Estimate.Permissions.GetByModelID(p.Source.(*estimate.Session).ID, paramSetFromGraphQLParams(util.KeyPermission, p, ctx.Logger)), nil
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
			Name: util.SvcEstimate.Title,
			Fields: graphql.Fields{
				util.KeyID: &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				util.KeySlug: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyTitle: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.WithID(util.SvcTeam.Key): &graphql.Field{
					Type: graphql.String,
				},
				util.WithID(util.SvcSprint.Key): &graphql.Field{
					Type: graphql.String,
				},
				util.KeyOwner: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyStatus: &graphql.Field{
					Type: graphql.NewNonNull(estimateStatusType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*estimate.Session).Status.Key, nil
					},
				},
				"choices": &graphql.Field{
					Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				util.Plural(util.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This estimate's members",
					Args:        listArgs,
					Resolve:     ctxF(estimateMemberResolver),
				},
			},
		},
	)
}
