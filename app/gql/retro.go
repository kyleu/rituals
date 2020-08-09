package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
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

	retroResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Retro(ctx.App).GetBySlug(npncore.MapGetString(p.Args, npncore.KeyKey, ctx.Logger)), nil
	}

	retrosResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Retro(ctx.App).List(paramSetFromGraphQLParams(svc.Key, params, ctx.Logger)), nil
	}

	retroActionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Action(ctx.App).GetBySvcModel(svc, p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(npncore.KeyAction, p, ctx.Logger)), nil
	}

	retroMemberResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Retro(ctx.App).Data.Members.GetByModelID(p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(npncore.KeyMember, p, ctx.Logger)), nil
	}

	retroPermissionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Retro(ctx.App).Data.Permissions.GetByModelID(p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(npncore.KeyPermission, p, ctx.Logger)), nil
	}

	retroCommentResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Retro(ctx.App).Data.GetComments(p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(npncore.KeyComment, p, ctx.Logger)), nil
	}

	retroHistoryResolver := func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		ret := app.Retro(ctx.App).Data.History.GetByModelID(p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(npncore.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	retroTeamResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		sess := p.Source.(*retro.Session)
		if sess.TeamID != nil {
			return app.Team(ctx.App).GetByID(*sess.TeamID), nil
		}
		return nil, nil
	}

	retroSprintResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		sess := p.Source.(*retro.Session)
		if sess.SprintID != nil {
			return app.Sprint(ctx.App).GetByID(*sess.SprintID), nil
		}
		return nil, nil
	}

	retroType = graphql.NewObject(
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
					Type: graphql.NewNonNull(retroStatusType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*retro.Session).Status.Key, nil
					},
				},
				npncore.Plural(npncore.KeyCategory): &graphql.Field{
					Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				npncore.Plural(npncore.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This retro's members",
					Args:        listArgs,
					Resolve:     ctxF(retroMemberResolver),
				},
				npncore.Plural(npncore.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This retro's comments",
					Args:        listArgs,
					Resolve:     ctxF(retroCommentResolver),
				},
				npncore.KeyHistory: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(historyType)),
					Description: "This retro's name history",
					Args:        listArgs,
					Resolve:     ctxF(retroHistoryResolver),
				},
			},
		},
	)
}
