package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	profileResolver         npngraphql.Callback
	profileTeamResolver     npngraphql.Callback
	profileSprintResolver   npngraphql.Callback
	profileEstimateResolver npngraphql.Callback
	profileStandupResolver  npngraphql.Callback
	profileRetroResolver    npngraphql.Callback
	profileType             *graphql.Object
)

func initProfile() {
	profileResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return ctx.Profile.ToProfile(), nil
	}

	profileTeamResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Team.GetByMember(p.Source.(npnuser.Profile).UserID, npngraphql.ParamSetFromGraphQLParams(util.SvcTeam.Key, p, ctx.Logger)), nil
	}

	profileSprintResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Sprint.GetByMember(p.Source.(npnuser.Profile).UserID, npngraphql.ParamSetFromGraphQLParams(util.SvcSprint.Key, p, ctx.Logger)), nil
	}

	profileEstimateResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Estimate.GetByMember(p.Source.(npnuser.Profile).UserID, npngraphql.ParamSetFromGraphQLParams(util.SvcEstimate.Key, p, ctx.Logger)), nil
	}

	profileStandupResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Standup.GetByMember(p.Source.(npnuser.Profile).UserID, npngraphql.ParamSetFromGraphQLParams(util.SvcStandup.Key, p, ctx.Logger)), nil
	}

	profileRetroResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Svc(ctx.App).Retro.GetByMember(p.Source.(npnuser.Profile).UserID, npngraphql.ParamSetFromGraphQLParams(util.SvcRetro.Key, p, ctx.Logger)), nil
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
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(profileTeamResolver),
				},
				util.SvcSprint.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(sprintType)),
					Description: "Your current sprints",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(profileSprintResolver),
				},
				util.SvcEstimate.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(estimateType)),
					Description: "Your current estimates",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(profileEstimateResolver),
				},
				util.SvcStandup.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(standupType)),
					Description: "Your current standups",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(profileStandupResolver),
				},
				util.SvcRetro.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(retroType)),
					Description: "Your current retros",
					Args:        npngraphql.ListArgs,
					Resolve:     npngraphql.CtxF(profileRetroResolver),
				},
			},
		},
	)

	memberType.AddFieldConfig(npncore.KeyUser, &graphql.Field{
		Type:        profileType,
		Description: "The user associated to this member",
		Resolve:     npngraphql.CtxF(memberProfileResolver),
	})

	commentType.AddFieldConfig(npncore.KeyUser, &graphql.Field{
		Type:        profileType,
		Description: "The user associated to this comment",
		Resolve:     npngraphql.CtxF(commentUserResolver),
	})

	actionType.AddFieldConfig(npncore.KeyUser, &graphql.Field{
		Type:        profileType,
		Description: "The user associated to this action",
		Resolve:     npngraphql.CtxF(actionUserResolver),
	})
}
