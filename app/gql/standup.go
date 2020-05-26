package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	standupArgs               graphql.FieldConfigArgument
	standupResolver           Callback
	standupsResolver          Callback
	standupMemberResolver     Callback
	standupPermissionResolver Callback
	standupTeamResolver       Callback
	standupSprintResolver     Callback
	standupType               *graphql.Object
)

func initStandup() {
	standupArgs = graphql.FieldConfigArgument{
		util.KeyKey: &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	standupStatusType := graphql.NewEnum(graphql.EnumConfig{
		Name: "StandupStatus",
		Values: graphql.EnumValueConfigMap{
			"new":     &graphql.EnumValueConfig{Value: "new"},
			"deleted": &graphql.EnumValueConfig{Value: "deleted"},
		},
	})

	standupResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		slug, err := paramKeyString(p)
		if err != nil {
			return nil, err
		}
		return ctx.App.Standup.GetBySlug(slug)
	}

	standupsResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Standup.List(paramSetFromGraphQLParams(util.SvcStandup.Key, params, ctx.Logger))
	}

	standupMemberResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Standup.Members.GetByModelID(p.Source.(*standup.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger)), nil
	}

	standupPermissionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Standup.Permissions.GetByModelID(p.Source.(*standup.Session).ID, paramSetFromGraphQLParams(util.KeyPermission, p, ctx.Logger)), nil
	}

	standupTeamResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		sess := p.Source.(*standup.Session)
		if sess.TeamID != nil {
			return ctx.App.Team.GetByID(*sess.TeamID)
		}
		return nil, nil
	}

	standupSprintResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		sess := p.Source.(*standup.Session)
		if sess.SprintID != nil {
			return ctx.App.Sprint.GetByID(*sess.SprintID)
		}
		return nil, nil
	}

	standupType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.SvcStandup.Title,
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
					Type: graphql.NewNonNull(standupStatusType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*standup.Session).Status.Key, nil
					},
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				util.Plural(util.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This standup's members",
					Args:        listArgs,
					Resolve:     ctxF(standupMemberResolver),
				},
			},
		},
	)
}
