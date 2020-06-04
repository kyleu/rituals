package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	retroResolver           Callback
	retrosResolver          Callback
	retroActionResolver     Callback
	retroPermissionResolver Callback
	retroTeamResolver       Callback
	retroSprintResolver     Callback
	retroType               *graphql.Object
)

func initRetro() {
	svc := util.SvcRetro

	retroStatusType := graphql.NewEnum(graphql.EnumConfig{
		Name: "RetroStatus",
		Values: graphql.EnumValueConfigMap{
			"new":     &graphql.EnumValueConfig{Value: "new"},
			"deleted": &graphql.EnumValueConfig{Value: "deleted"},
		},
	})

	retroResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.GetBySlug(util.MapGetString(p.Args, util.KeyKey, ctx.Logger)), nil
	}

	retrosResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.List(paramSetFromGraphQLParams(svc.Key, params, ctx.Logger)), nil
	}

	retroActionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Action.GetBySvcModel(svc, p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(util.KeyAction, p, ctx.Logger)), nil
	}

	retroMemberResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.Data.Members.GetByModelID(p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger)), nil
	}

	retroPermissionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.Data.Permissions.GetByModelID(p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(util.KeyPermission, p, ctx.Logger)), nil
	}

	retroCommentResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.Data.GetComments(p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(util.KeyComment, p, ctx.Logger)), nil
	}

	retroHistoryResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		ret := ctx.App.Retro.Data.History.GetByModelID(p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(util.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	retroTeamResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		sess := p.Source.(*retro.Session)
		if sess.TeamID != nil {
			return ctx.App.Team.GetByID(*sess.TeamID), nil
		}
		return nil, nil
	}

	retroSprintResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		sess := p.Source.(*retro.Session)
		if sess.SprintID != nil {
			return ctx.App.Sprint.GetByID(*sess.SprintID), nil
		}
		return nil, nil
	}

	retroType = graphql.NewObject(
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
					Type: graphql.NewNonNull(retroStatusType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*retro.Session).Status.Key, nil
					},
				},
				util.Plural(util.KeyCategory): &graphql.Field{
					Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				util.Plural(util.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This retro's members",
					Args:        listArgs,
					Resolve:     ctxF(retroMemberResolver),
				},
				util.Plural(util.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This retro's comments",
					Args:        listArgs,
					Resolve:     ctxF(retroCommentResolver),
				},
				util.KeyHistory: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(historyType)),
					Description: "This retro's name history",
					Args:        listArgs,
					Resolve:     ctxF(retroHistoryResolver),
				},
			},
		},
	)
}
