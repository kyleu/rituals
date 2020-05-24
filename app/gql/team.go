package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	teamArgs           graphql.FieldConfigArgument
	teamResolver       Callback
	teamsResolver      Callback
	teamMemberResolver Callback
	teamType           *graphql.Object
)

func initTeam() {
	teamArgs = graphql.FieldConfigArgument{
		util.KeyKey: &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	teamResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		slug, err := paramString(p, util.KeyKey)
		if err != nil {
			return nil, err
		}
		return ctx.App.Team.GetBySlug(slug)
	}

	teamsResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Team.List(paramSetFromGraphQLParams(util.SvcTeam.Key, params, ctx.Logger))
	}

	teamMemberResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Team.Members.GetByModelID(p.Source.(*team.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger)), nil
	}

	teamType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Team",
			Fields: graphql.Fields{
				util.KeyID: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"slug": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"title": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"owner": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				util.KeyPlural(util.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This team's members",
					Args:        listArgs,
					Resolve:     ctxF(teamMemberResolver),
				},
			},
		},
	)

	estimateType.AddFieldConfig(util.SvcTeam.Key, &graphql.Field{
		Type:        teamType,
		Description: "This estimate's team",
		Args:        listArgs,
		Resolve:     ctxF(estimateTeamResolver),
	})

	standupType.AddFieldConfig(util.SvcTeam.Key, &graphql.Field{
		Type:        teamType,
		Description: "This standup's team",
		Args:        listArgs,
		Resolve:     ctxF(standupTeamResolver),
	})

	retroType.AddFieldConfig(util.SvcTeam.Key, &graphql.Field{
		Type:        teamType,
		Description: "This retro's team",
		Args:        listArgs,
		Resolve:     ctxF(retroTeamResolver),
	})

	sprintType.AddFieldConfig(util.SvcTeam.Key, &graphql.Field{
		Type:        teamType,
		Description: "This sprint's team",
		Args:        listArgs,
		Resolve:     ctxF(sprintTeamResolver),
	})
}
