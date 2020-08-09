package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	sprintResolver           Callback
	sprintsResolver          Callback
	sprintActionResolver     Callback
	sprintPermissionResolver Callback
	sprintTeamResolver       Callback
	sprintType               *graphql.Object
)

func initSprint() {
	svc := util.SvcSprint

	sprintResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Sprint(ctx.App).GetBySlug(npncore.MapGetString(p.Args, npncore.KeyKey, ctx.Logger)), nil
	}

	sprintsResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Sprint(ctx.App).List(paramSetFromGraphQLParams(svc.Key, params, ctx.Logger)), nil
	}

	sprintTeamResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		sess := p.Source.(*sprint.Session)
		if sess.TeamID != nil {
			return app.Team(ctx.App).GetByID(*sess.TeamID), nil
		}
		return nil, nil
	}

	sprintActionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Action(ctx.App).GetBySvcModel(svc, p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(npncore.KeyAction, p, ctx.Logger)), nil
	}

	sprintMemberResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Sprint(ctx.App).Data.Members.GetByModelID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(npncore.KeyMember, p, ctx.Logger)), nil
	}

	sprintPermissionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Sprint(ctx.App).Data.Permissions.GetByModelID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(npncore.KeyPermission, p, ctx.Logger)), nil
	}

	sprintCommentResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Sprint(ctx.App).Data.GetComments(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(npncore.KeyComment, p, ctx.Logger)), nil
	}

	sprintHistoryResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		ret := app.Sprint(ctx.App).Data.History.GetByModelID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(npncore.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	sprintEstimateResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Estimate(ctx.App).GetBySprintID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.SvcEstimate.Key, p, ctx.Logger)), nil
	}

	sprintStandupResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Standup(ctx.App).GetBySprintID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.SvcStandup.Key, p, ctx.Logger)), nil
	}

	sprintRetroResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Retro(ctx.App).GetBySprintID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.SvcRetro.Key, p, ctx.Logger)), nil
	}

	sprintType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: svc.Title,
			Fields: graphql.Fields{
				npncore.KeyID: &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
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
					Args:        listArgs,
					Resolve:     ctxF(sprintMemberResolver),
				},
				util.SvcEstimate.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(estimateType)),
					Description: "This sprint's estimates",
					Args:        listArgs,
					Resolve:     ctxF(sprintEstimateResolver),
				},
				util.SvcStandup.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(standupType)),
					Description: "This sprint's standups",
					Args:        listArgs,
					Resolve:     ctxF(sprintStandupResolver),
				},
				util.SvcRetro.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(retroType)),
					Description: "This sprint's retros",
					Args:        listArgs,
					Resolve:     ctxF(sprintRetroResolver),
				},
				npncore.Plural(npncore.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This sprint's comments",
					Args:        listArgs,
					Resolve:     ctxF(sprintCommentResolver),
				},
				npncore.KeyHistory: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(historyType)),
					Description: "This sprint's name history",
					Args:        listArgs,
					Resolve:     ctxF(sprintHistoryResolver),
				},
			},
		},
	)

	estimateType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        sprintType,
		Description: "This estimate's sprint",
		Args:        listArgs,
		Resolve:     ctxF(estimateSprintResolver),
	})

	standupType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        sprintType,
		Description: "This standup's sprint",
		Args:        listArgs,
		Resolve:     ctxF(standupSprintResolver),
	})

	retroType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        sprintType,
		Description: "This retro's sprint",
		Args:        listArgs,
		Resolve:     ctxF(retroSprintResolver),
	})
}
