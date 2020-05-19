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
	sprintMemberResolver   Callback
	sprintEstimateResolver Callback
	sprintStandupResolver  Callback
	sprintRetroResolver    Callback
	sprintType             *graphql.Object
)

func initSprint() {
	sprintArgs = graphql.FieldConfigArgument{
		"key": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}

	sprintResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		slug, ok := p.Args["key"].(string)
		if ok {
			return ctx.App.Sprint.GetBySlug(slug)
		}
		return nil, nil
	}

	sprintsResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.List(paramSetFromGraphQLParams(util.SvcSprint.Key, params, ctx.Logger))
	}

	sprintMemberResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.Members.GetByModelID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger))
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
				"id": &graphql.Field{
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
				"startDate": &graphql.Field{
					Type: graphql.String,
				},
				"endDate": &graphql.Field{
					Type: graphql.String,
				},
				"created": &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"members": &graphql.Field{
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
