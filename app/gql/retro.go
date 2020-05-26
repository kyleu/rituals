package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	retroArgs               graphql.FieldConfigArgument
	retroResolver           Callback
	retrosResolver          Callback
	retroMemberResolver     Callback
	retroPermissionResolver Callback
	retroTeamResolver       Callback
	retroSprintResolver     Callback
	retroType               *graphql.Object
)

func initRetro() {
	retroArgs = graphql.FieldConfigArgument{
		util.KeyKey: &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	retroStatusType := graphql.NewEnum(graphql.EnumConfig{
		Name: "RetroStatus",
		Values: graphql.EnumValueConfigMap{
			"new":     &graphql.EnumValueConfig{Value: "new"},
			"deleted": &graphql.EnumValueConfig{Value: "deleted"},
		},
	})

	retroResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		slug, err := paramKeyString(p)
		if err != nil {
			return nil, err
		}
		return ctx.App.Retro.GetBySlug(slug)
	}

	retrosResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.List(paramSetFromGraphQLParams(util.SvcRetro.Key, params, ctx.Logger))
	}

	retroMemberResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.Members.GetByModelID(p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger)), nil
	}

	retroPermissionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.Permissions.GetByModelID(p.Source.(*retro.Session).ID, paramSetFromGraphQLParams(util.KeyPermission, p, ctx.Logger)), nil
	}

	retroTeamResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		sess := p.Source.(*retro.Session)
		if sess.TeamID != nil {
			return ctx.App.Team.GetByID(*sess.TeamID)
		}
		return nil, nil
	}

	retroSprintResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		sess := p.Source.(*retro.Session)
		if sess.SprintID != nil {
			return ctx.App.Sprint.GetByID(*sess.SprintID)
		}
		return nil, nil
	}

	retroType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.SvcRetro.Title,
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
				"categories": &graphql.Field{
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
			},
		},
	)
}
