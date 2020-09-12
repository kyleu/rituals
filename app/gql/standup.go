package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	standupResolver           npngraphql.Callback
	standupsResolver          npngraphql.Callback
	standupActionResolver     npngraphql.Callback
	standupPermissionResolver npngraphql.Callback
	standupTeamResolver       npngraphql.Callback
	standupSprintResolver     npngraphql.Callback
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
		return app.Svc(ctx.App).Standup.GetBySlug(npncore.MapGetString(p.Args, npncore.KeyKey, ctx.Logger)), nil
	}

	standupsResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Standup.List(npngraphql.ParamSetFromGraphQLParams(svc.Key, params, ctx.Logger)), nil
	}

	standupActionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Action.GetBySvcModel(svc.Key, p.Source.(*standup.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyAction, p, ctx.Logger)), nil
	}

	standupMemberResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Standup.Data.Members.GetByModelID(p.Source.(*standup.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyMember, p, ctx.Logger)), nil
	}

	standupPermissionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Standup.Data.Permissions.GetByModelID(p.Source.(*standup.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyPermission, p, ctx.Logger)), nil
	}

	standupCommentResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Standup.Data.GetComments(p.Source.(*standup.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyComment, p, ctx.Logger)), nil
	}

	standupHistoryResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		ret := app.Svc(ctx.App).Standup.Data.History.GetByModelID(p.Source.(*standup.Session).ID, npngraphql.ParamSetFromGraphQLParams(npncore.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	standupTeamResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		sess := p.Source.(*standup.Session)
		if sess.TeamID != nil {
			return app.Svc(ctx.App).Team.GetByID(*sess.TeamID), nil
		}
		return nil, nil
	}

	standupSprintResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		sess := p.Source.(*standup.Session)
		if sess.SprintID != nil {
			return app.Svc(ctx.App).Sprint.GetByID(*sess.SprintID), nil
		}
		return nil, nil
	}

	standupType = graphql.NewObject(
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
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(standupMemberResolver),
				},
				npncore.Plural(npncore.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This standups's comments",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(standupCommentResolver),
				},
				npncore.KeyHistory: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(historyType)),
					Description: "This standup's name history",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(standupHistoryResolver),
				},
			},
		},
	)
}
