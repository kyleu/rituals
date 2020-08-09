package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
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

	teamResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).GetBySlug(npncore.MapGetString(p.Args, npncore.KeyKey, ctx.Logger)), nil
	}

	teamsResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).List(paramSetFromGraphQLParams(svc.Key, params, ctx.Logger)), nil
	}

	teamActionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Action(ctx.App).GetBySvcModel(svc, p.Source.(*team.Session).ID, paramSetFromGraphQLParams(npncore.KeyAction, p, ctx.Logger)), nil
	}

	teamMemberResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).Data.Members.GetByModelID(p.Source.(*team.Session).ID, paramSetFromGraphQLParams(npncore.KeyMember, p, ctx.Logger)), nil
	}

	teamPermissionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).Data.Permissions.GetByModelID(p.Source.(*team.Session).ID, paramSetFromGraphQLParams(npncore.KeyPermission, p, ctx.Logger)), nil
	}

	teamCommentResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).Data.GetComments(p.Source.(*team.Session).ID, paramSetFromGraphQLParams(npncore.KeyComment, p, ctx.Logger)), nil
	}

	teamHistoryResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		ret := app.Team(ctx.App).Data.History.GetByModelID(p.Source.(*team.Session).ID, paramSetFromGraphQLParams(npncore.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	teamType = graphql.NewObject(
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
				npncore.KeyOwner: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				npncore.Plural(npncore.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This team's members",
					Args:        listArgs,
					Resolve:     ctxF(teamMemberResolver),
				},
				npncore.Plural(npncore.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This team's comments",
					Args:        listArgs,
					Resolve:     ctxF(teamCommentResolver),
				},
				npncore.KeyHistory: &graphql.Field{
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
