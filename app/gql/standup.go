package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	standupResolver           Callback
	standupsResolver          Callback
	standupActionResolver     Callback
	standupPermissionResolver Callback
	standupTeamResolver       Callback
	standupSprintResolver     Callback
	standupType               *graphql.Object
)

func initStandup() {
	svc := util.SvcStandup

	standupStatusType := graphql.NewEnum(graphql.EnumConfig{
		Name: "StandupStatus",
		Values: graphql.EnumValueConfigMap{
			"new":     &graphql.EnumValueConfig{Value: "new"},
			"deleted": &graphql.EnumValueConfig{Value: "deleted"},
		},
	})

	standupResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Standup(ctx.App).GetBySlug(npncore.MapGetString(p.Args, npncore.KeyKey, ctx.Logger)), nil
	}

	standupsResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Standup(ctx.App).List(paramSetFromGraphQLParams(svc.Key, params, ctx.Logger)), nil
	}

	standupActionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Action(ctx.App).GetBySvcModel(svc, p.Source.(*standup.Session).ID, paramSetFromGraphQLParams(npncore.KeyAction, p, ctx.Logger)), nil
	}

	standupMemberResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Standup(ctx.App).Data.Members.GetByModelID(p.Source.(*standup.Session).ID, paramSetFromGraphQLParams(npncore.KeyMember, p, ctx.Logger)), nil
	}

	standupPermissionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Standup(ctx.App).Data.Permissions.GetByModelID(p.Source.(*standup.Session).ID, paramSetFromGraphQLParams(npncore.KeyPermission, p, ctx.Logger)), nil
	}

	standupCommentResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Standup(ctx.App).Data.GetComments(p.Source.(*standup.Session).ID, paramSetFromGraphQLParams(npncore.KeyComment, p, ctx.Logger)), nil
	}

	standupHistoryResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		ret := app.Standup(ctx.App).Data.History.GetByModelID(p.Source.(*standup.Session).ID, paramSetFromGraphQLParams(npncore.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	standupTeamResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		sess := p.Source.(*standup.Session)
		if sess.TeamID != nil {
			return app.Team(ctx.App).GetByID(*sess.TeamID), nil
		}
		return nil, nil
	}

	standupSprintResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		sess := p.Source.(*standup.Session)
		if sess.SprintID != nil {
			return app.Sprint(ctx.App).GetByID(*sess.SprintID), nil
		}
		return nil, nil
	}

	standupType = graphql.NewObject(
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
				npncore.WithID(util.SvcSprint.Key): &graphql.Field{
					Type: graphql.String,
				},
				npncore.KeyOwner: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyStatus: &graphql.Field{
					Type: graphql.NewNonNull(standupStatusType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*standup.Session).Status.Key, nil
					},
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				npncore.Plural(npncore.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This standup's members",
					Args:        listArgs,
					Resolve:     ctxF(standupMemberResolver),
				},
				npncore.Plural(npncore.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This standups's comments",
					Args:        listArgs,
					Resolve:     ctxF(standupCommentResolver),
				},
				npncore.KeyHistory: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(historyType)),
					Description: "This standup's name history",
					Args:        listArgs,
					Resolve:     ctxF(standupHistoryResolver),
				},
			},
		},
	)
}
