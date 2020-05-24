package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	sprintArgs             graphql.FieldConfigArgument
	sprintResolver         Callback
	sprintsResolver        Callback
	sprintTeamResolver     Callback
	sprintMemberResolver   Callback
	sprintEstimateResolver Callback
	sprintStandupResolver  Callback
	sprintRetroResolver    Callback
	sprintType             *graphql.Object
)

func initSprint() {
	sprintArgs = graphql.FieldConfigArgument{
		util.KeyKey: &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	sprintResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		slug, err := paramString(p, util.KeyKey)
		if err != nil {
			return nil, err
		}
		return ctx.App.Sprint.GetBySlug(slug)
	}

	sprintsResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.List(paramSetFromGraphQLParams(util.SvcSprint.Key, params, ctx.Logger))
	}

	sprintTeamResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		sess := p.Source.(*sprint.Session)
		if sess.TeamID != nil {
			return ctx.App.Team.GetByID(*sess.TeamID)
		}
		return nil, nil
	}

	sprintMemberResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.Members.GetByModelID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger)), nil
	}

	sprintEstimateResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Estimate.GetBySprint(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.SvcEstimate.Key, p, ctx.Logger))
	}

	sprintStandupResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Standup.GetBySprint(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.SvcStandup.Key, p, ctx.Logger))
	}

	sprintRetroResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.GetBySprint(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.SvcRetro.Key, p, ctx.Logger))
	}

	sprintType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Sprint",
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
				"teamID": &graphql.Field{
					Type: graphql.String,
				},
				"owner": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"startDate": &graphql.Field{
					Type: graphql.String,
				},
				"endDate": &graphql.Field{
					Type: graphql.String,
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				util.KeyPlural(util.KeyMember): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(memberType)),
					Description: "This sprint's members",
					Args:        listArgs,
					Resolve:     ctxF(sprintMemberResolver),
				},
				util.SvcEstimate.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(estimateType)),
					Description: "This sprint's estimates",
					Args:        listArgs,
					Resolve:     ctxF(sprintEstimateResolver),
				},
				util.SvcStandup.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(standupType)),
					Description: "This sprint's standups",
					Args:        listArgs,
					Resolve:     ctxF(sprintStandupResolver),
				},
				util.SvcRetro.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(retroType)),
					Description: "This sprint's retros",
					Args:        listArgs,
					Resolve:     ctxF(sprintRetroResolver),
				},
			},
		},
	)

	estimateType.AddFieldConfig(util.SvcSprint.Key, &graphql.Field{
		Type:        sprintType,
		Description: "This estimate's sprint",
		Args:        listArgs,
		Resolve:     ctxF(estimateSprintResolver),
	})

	standupType.AddFieldConfig(util.SvcSprint.Key, &graphql.Field{
		Type:        sprintType,
		Description: "This standup's sprint",
		Args:        listArgs,
		Resolve:     ctxF(standupSprintResolver),
	})

	retroType.AddFieldConfig(util.SvcSprint.Key, &graphql.Field{
		Type:        sprintType,
		Description: "This retro's sprint",
		Args:        listArgs,
		Resolve:     ctxF(retroSprintResolver),
	})
}
