package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	sprintResolver           npngraphql.Callback
	sprintsResolver          npngraphql.Callback
	sprintActionResolver     npngraphql.Callback
	sprintPermissionResolver npngraphql.Callback
	sprintTeamResolver       npngraphql.Callback
	sprintType               *graphql.Object
)

func initSprint() {
	svc := util.SvcSprint

	sprintResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Sprint.GetBySlug(npncore.MapGetString(p.Args, npncore.KeyKey, ctx.Logger)), nil
	}

	sprintsResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Sprint.List(npngraphql.ParamSetFromGraphQLParams(svc.Key, params, ctx.Logger)), nil
	}

	sprintTeamResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		sess := p.Source.(*sprint.Session)
		if sess.TeamID != nil {
			return app.Svc(ctx.App).Team.GetByID(*sess.TeamID), nil
		}
		return nil, nil
	}

	sprintActionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Action.GetBySvcModel(svc.Key, p.Source.(*sprint.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyAction, p, ctx.Logger)), nil
	}

	sprintMemberResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Sprint.Data.Members.GetByModelID(p.Source.(*sprint.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyMember, p, ctx.Logger)), nil
	}

	sprintPermissionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Sprint.Data.Permissions.GetByModelID(p.Source.(*sprint.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyPermission, p, ctx.Logger)), nil
	}

	sprintCommentResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Sprint.Data.GetComments(p.Source.(*sprint.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyComment, p, ctx.Logger)), nil
	}

	sprintHistoryResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		ret := app.Svc(ctx.App).Sprint.Data.History.GetByModelID(p.Source.(*sprint.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	sprintEstimateResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Estimate.GetBySprintID(p.Source.(*sprint.Session).ID, npngraphql.ParamSetFromGraphQLParams(util.SvcEstimate.Key, p, ctx.Logger)), nil
	}

	sprintStandupResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Standup.GetBySprintID(p.Source.(*sprint.Session).ID, npngraphql.ParamSetFromGraphQLParams(util.SvcStandup.Key, p, ctx.Logger)), nil
	}

	sprintRetroResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Retro.GetBySprintID(p.Source.(*sprint.Session).ID, npngraphql.ParamSetFromGraphQLParams(util.SvcRetro.Key, p, ctx.Logger)), nil
	}

	sprintType = graphql.NewObject(
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
				npncore.KeyOwner: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"startDate": &graphql.Field{
					Type: graphql.String,
				},
				"endDate": &graphql.Field{
					Type: graphql.String,
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				npncore.Plural(npncore.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This sprint's members",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(sprintMemberResolver),
				},
				util.SvcEstimate.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(estimateType)),
					Description: "This sprint's estimates",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(sprintEstimateResolver),
				},
				util.SvcStandup.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(standupType)),
					Description: "This sprint's standups",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(sprintStandupResolver),
				},
				util.SvcRetro.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(retroType)),
					Description: "This sprint's retros",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(sprintRetroResolver),
				},
				npncore.Plural(npncore.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This sprint's comments",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(sprintCommentResolver),
				},
				npncore.KeyHistory: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(historyType)),
					Description: "This sprint's name history",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(sprintHistoryResolver),
				},
			},
		},
	)

	estimateType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        sprintType,
		Description: "This estimate's sprint",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(estimateSprintResolver),
	})

	standupType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        sprintType,
		Description: "This standup's sprint",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(standupSprintResolver),
	})

	retroType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        sprintType,
		Description: "This retro's sprint",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(retroSprintResolver),
	})
}
