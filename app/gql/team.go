package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	teamArgs               graphql.FieldConfigArgument
	teamResolver           Callback
	teamsResolver          Callback
	teamMemberResolver     Callback
	teamPermissionResolver Callback
	teamType               *graphql.Object
)

func initTeam() {
	teamArgs = graphql.FieldConfigArgument{
		util.KeyKey: &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	teamResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		slug, err := paramKeyString(p)
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

	teamPermissionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Team.Permissions.GetByModelID(p.Source.(*team.Session).ID, paramSetFromGraphQLParams(util.KeyPermission, p, ctx.Logger)), nil
	}

	teamType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.SvcTeam.Title,
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
