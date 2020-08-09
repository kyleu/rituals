package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	profileResolver         Callback
	profileTeamResolver     Callback
	profileSprintResolver   Callback
	profileEstimateResolver Callback
	profileStandupResolver  Callback
	profileRetroResolver    Callback
	profileType             *graphql.Object
)

func initProfile() {
	profileResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return ctx.Profile.ToProfile(), nil
	}

	profileTeamResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Team(ctx.App).GetByMember(p.Source.(npnuser.Profile).UserID, paramSetFromGraphQLParams(util.SvcTeam.Key, p, ctx.Logger)), nil
	}

	profileSprintResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Sprint(ctx.App).GetByMember(p.Source.(npnuser.Profile).UserID, paramSetFromGraphQLParams(util.SvcSprint.Key, p, ctx.Logger)), nil
	}

	profileEstimateResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Estimate(ctx.App).GetByMember(p.Source.(npnuser.Profile).UserID, paramSetFromGraphQLParams(util.SvcEstimate.Key, p, ctx.Logger)), nil
	}

	profileStandupResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Standup(ctx.App).GetByMember(p.Source.(npnuser.Profile).UserID, paramSetFromGraphQLParams(util.SvcStandup.Key, p, ctx.Logger)), nil
	}

	profileRetroResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Retro(ctx.App).GetByMember(p.Source.(npnuser.Profile).UserID, paramSetFromGraphQLParams(util.SvcRetro.Key, p, ctx.Logger)), nil
	}

	profileType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Profile",
			Fields: graphql.Fields{
				npncore.WithID(npncore.KeyUser): &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyName: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyRole: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyTheme: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"navColor": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"linkColor": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"picture": &graphql.Field{
					Type: graphql.String,
				},
				"locale": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.SvcTeam.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(teamType)),
					Description: "Your current teams",
					Args:        listArgs,
					Resolve:     ctxF(profileTeamResolver),
				},
				util.SvcSprint.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(sprintType)),
					Description: "Your current sprints",
					Args:        listArgs,
					Resolve:     ctxF(profileSprintResolver),
				},
				util.SvcEstimate.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(estimateType)),
					Description: "Your current estimates",
					Args:        listArgs,
					Resolve:     ctxF(profileEstimateResolver),
				},
				util.SvcStandup.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(standupType)),
					Description: "Your current standups",
					Args:        listArgs,
					Resolve:     ctxF(profileStandupResolver),
				},
				util.SvcRetro.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(retroType)),
					Description: "Your current retros",
					Args:        listArgs,
					Resolve:     ctxF(profileRetroResolver),
				},
			},
		},
	)

	memberType.AddFieldConfig(npncore.KeyUser, &graphql.Field{
		Type:        profileType,
		Description: "The user associated to this member",
		Resolve:     ctxF(memberProfileResolver),
	})

	commentType.AddFieldConfig(npncore.KeyUser, &graphql.Field{
		Type:        profileType,
		Description: "The user associated to this comment",
		Resolve:     ctxF(commentUserResolver),
	})

	actionType.AddFieldConfig(npncore.KeyUser, &graphql.Field{
		Type:        profileType,
		Description: "The user associated to this action",
		Resolve:     ctxF(actionUserResolver),
	})
}
