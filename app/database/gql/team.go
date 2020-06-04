package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	teamResolver           Callback
	teamsResolver          Callback
	teamActionResolver     Callback
	teamPermissionResolver Callback
	teamType               *graphql.Object
)

func initTeam() {
	svc := util.SvcTeam

	teamResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Team.GetBySlug(util.MapGetString(p.Args, util.KeyKey, ctx.Logger)), nil
	}

	teamsResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Team.List(paramSetFromGraphQLParams(svc.Key, params, ctx.Logger)), nil
	}

	teamActionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Action.GetBySvcModel(svc, p.Source.(*team.Session).ID, paramSetFromGraphQLParams(util.KeyAction, p, ctx.Logger)), nil
	}

	teamMemberResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Team.Data.Members.GetByModelID(p.Source.(*team.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger)), nil
	}

	teamPermissionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Team.Data.Permissions.GetByModelID(p.Source.(*team.Session).ID, paramSetFromGraphQLParams(util.KeyPermission, p, ctx.Logger)), nil
	}

	teamCommentResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Team.Data.GetComments(p.Source.(*team.Session).ID, paramSetFromGraphQLParams(util.KeyComment, p, ctx.Logger)), nil
	}

	teamHistoryResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		ret := ctx.App.Team.Data.History.GetByModelID(p.Source.(*team.Session).ID, paramSetFromGraphQLParams(util.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	teamType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: svc.Title,
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
				util.KeyOwner: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				util.Plural(util.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This team's members",
					Args:        listArgs,
					Resolve:     ctxF(teamMemberResolver),
				},
				util.Plural(util.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This team's comments",
					Args:        listArgs,
					Resolve:     ctxF(teamCommentResolver),
				},
				util.KeyHistory: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(historyType)),
					Description: "This team's name history",
					Args:        listArgs,
					Resolve:     ctxF(teamHistoryResolver),
				},
			},
		},
	)

	sprintType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        teamType,
		Description: "This sprint's team",
		Args:        listArgs,
		Resolve:     ctxF(sprintTeamResolver),
	})

	estimateType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        teamType,
		Description: "This estimate's team",
		Args:        listArgs,
		Resolve:     ctxF(estimateTeamResolver),
	})

	standupType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        teamType,
		Description: "This standup's team",
		Args:        listArgs,
		Resolve:     ctxF(standupTeamResolver),
	})

	retroType.AddFieldConfig(svc.Key, &graphql.Field{
		Type:        teamType,
		Description: "This retro's team",
		Args:        listArgs,
		Resolve:     ctxF(retroTeamResolver),
	})
}
