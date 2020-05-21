package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	standupArgs           graphql.FieldConfigArgument
	standupResolver       Callback
	standupsResolver      Callback
	standupTeamResolver   Callback
	standupSprintResolver Callback
	standupMemberResolver Callback
	standupType           *graphql.Object
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
		slug, ok := p.Args[util.KeyKey]
		if ok {
			return ctx.App.Standup.GetBySlug(slug.(string))
		}
		return nil, nil
	}

	standupsResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Standup.List(paramSetFromGraphQLParams(util.SvcStandup.Key, params, ctx.Logger))
	}

	standupMemberResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Standup.Members.GetByModelID(p.Source.(*standup.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger)), nil
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
			Name: "Standup",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"slug": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"title": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"teamID": &graphql.Field{
					Type: graphql.String,
				},
				"sprintID": &graphql.Field{
					Type: graphql.String,
				},
				"owner": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"status": &graphql.Field{
					Type: graphql.NewNonNull(standupStatusType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*standup.Session).Status.Key, nil
					},
				},
				"created": &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"members": &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This standup's members",
					Args:        listArgs,
					Resolve:     ctxF(standupMemberResolver),
				},
			},
		},
	)
}
