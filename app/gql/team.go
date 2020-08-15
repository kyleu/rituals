package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	teamResolver           npngraphql.Callback
	teamsResolver          npngraphql.Callback
	teamActionResolver     npngraphql.Callback
	teamPermissionResolver npngraphql.Callback
	teamType               *graphql.Object
)

func initTeam() {
	svc := util.SvcTeam

	teamResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).GetBySlug(npncore.MapGetString(p.Args, npncore.KeyKey, ctx.Logger)), nil
	}

	teamsResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).List(npngraphql.ParamSetFromGraphQLParams(svc.Key, params, ctx.Logger)), nil
	}

	teamActionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Action(ctx.App).GetBySvcModel(svc.Key, p.Source.(*team.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyAction, p, ctx.Logger)), nil
	}

	teamMemberResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).Data.Members.GetByModelID(p.Source.(*team.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyMember, p, ctx.Logger)), nil
	}

	teamPermissionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).Data.Permissions.GetByModelID(p.Source.(*team.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyPermission, p, ctx.Logger)), nil
	}

	teamCommentResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).Data.GetComments(p.Source.(*team.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyComment, p, ctx.Logger)), nil
	}

	teamHistoryResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		ret := app.Team(ctx.App).Data.History.GetByModelID(p.Source.(*team.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	teamType = graphql.NewObject(
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
				npncore.KeyOwner: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				npncore.Plural(npncore.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This team's members",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(teamMemberResolver),
				},
				npncore.Plural(npncore.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This team's comments",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(teamCommentResolver),
				},
				npncore.KeyHistory: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(historyType)),
					Description: "This team's name history",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(teamHistoryResolver),
				},
			},
		},
	)

	sprintType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        teamType,
		Description: "This sprint's team",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(sprintTeamResolver),
	})

	estimateType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        teamType,
		Description: "This estimate's team",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(estimateTeamResolver),
	})

	standupType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        teamType,
		Description: "This standup's team",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(standupTeamResolver),
	})

	retroType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        teamType,
		Description: "This retro's team",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(retroTeamResolver),
	})
}
