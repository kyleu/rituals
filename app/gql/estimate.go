package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	estimateResolver           npngraphql.Callback
	estimatesResolver          npngraphql.Callback
	estimateActionResolver     npngraphql.Callback
	estimatePermissionResolver npngraphql.Callback
	estimateTeamResolver       npngraphql.Callback
	estimateSprintResolver     npngraphql.Callback
	estimateType               *graphql.Object
)

func initEstimate() {
	svc := util.SvcEstimate
	initStory()

	estimateStatusType := graphql.NewEnum(graphql.EnumConfig{
		Name: "EstimateStatus",
		Values: graphql.EnumValueConfigMap{
			"new":      &graphql.EnumValueConfig{Value: "new"},
			"active":   &graphql.EnumValueConfig{Value: "active"},
			"complete": &graphql.EnumValueConfig{Value: "complete"},
			"deleted":  &graphql.EnumValueConfig{Value: "deleted"},
		},
	})

	estimateResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Estimate.GetBySlug(npncore.MapGetString(p.Args, npncore.KeyKey, ctx.Logger)), nil
	}

	estimatesResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Estimate.List(npngraphql.ParamSetFromGraphQLParams(svc.Key, params, ctx.Logger)), nil
	}

	estimateActionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Action.GetBySvcModel(svc.Key, p.Source.(*estimate.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyAction, p, ctx.Logger)), nil
	}

	estimatePermissionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Estimate.Data.Permissions.GetByModelID(p.Source.(*estimate.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyPermission, p, ctx.Logger)), nil
	}

	estimateMemberResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Estimate.Data.Members.GetByModelID(p.Source.(*estimate.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyMember, p, ctx.Logger)), nil
	}

	estimateCommentResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Estimate.Data.GetComments(p.Source.(*estimate.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyComment, p, ctx.Logger)), nil
	}

	estimateHistoryResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		ret := app.Svc(ctx.App).Estimate.Data.History.GetByModelID(p.Source.(*estimate.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	estimateTeamResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		sess := p.Source.(*estimate.Session)
		if sess.TeamID != nil {
			return app.Svc(ctx.App).Team.GetByID(*sess.TeamID), nil
		}
		return nil, nil
	}

	estimateSprintResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		sess := p.Source.(*estimate.Session)
		if sess.SprintID != nil {
			return app.Svc(ctx.App).Sprint.GetByID(*sess.SprintID), nil
		}
		return nil, nil
	}

	estimateType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: svc.Title,
			Fields: graphql.Fields{
				npncore.KeyID: &graphql.Field{
					Type: graphql.NewNonNull(npngraphql.ScalarUUID),
				},
				npncore.KeySlug: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyTitle: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.WithID(util.SvcTeam.Key): &graphql.Field{
					Type: graphql.String,
				},
				npncore.WithID(util.SvcSprint.Key): &graphql.Field{
					Type: graphql.String,
				},
				npncore.KeyOwner: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyStatus: &graphql.Field{
					Type: graphql.NewNonNull(estimateStatusType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*estimate.Session).Status.Key, nil
					},
				},
				npncore.Plural(npncore.KeyChoice): &graphql.Field{
					Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				npncore.Plural(util.KeyStory): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(storyType)),
					Description: "This estimate's stories",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(storiesResolver),
				},
				npncore.Plural(npncore.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This estimate's members",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(estimateMemberResolver),
				},
				npncore.Plural(npncore.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This estimate's comments",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(estimateCommentResolver),
				},
				npncore.KeyHistory: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(historyType)),
					Description: "This estimate's name history",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(estimateHistoryResolver),
				},
			},
		},
	)
}
