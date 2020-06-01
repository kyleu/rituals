package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	sprintResolver           Callback
	sprintsResolver          Callback
	sprintActionResolver     Callback
	sprintPermissionResolver Callback
	sprintTeamResolver       Callback
	sprintType               *graphql.Object
)

func initSprint() {
	sprintResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.GetBySlug(util.MapGetString(p.Args, util.KeyKey, ctx.Logger)), nil
	}

	sprintsResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.List(paramSetFromGraphQLParams(util.SvcSprint.Key, params, ctx.Logger)), nil
	}

	sprintTeamResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		sess := p.Source.(*sprint.Session)
		if sess.TeamID != nil {
			return ctx.App.Team.GetByID(*sess.TeamID), nil
		}
		return nil, nil
	}

	sprintActionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Action.GetBySvcModel(util.SvcSprint, p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.KeyAction, p, ctx.Logger)), nil
	}

	sprintMemberResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.Data.Members.GetByModelID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.KeyMember, p, ctx.Logger)), nil
	}

	sprintPermissionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.Data.Permissions.GetByModelID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.KeyPermission, p, ctx.Logger)), nil
	}

	sprintCommentResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.Data.Comments.GetByModelID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.KeyComment, p, ctx.Logger)), nil
	}


	sprintHistoryResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		ret := ctx.App.Sprint.Data.History.GetByModelID(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.KeyHistory, p, ctx.Logger))
		return ret, nil
	}

	sprintEstimateResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Estimate.GetBySprint(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.SvcEstimate.Key, p, ctx.Logger)), nil
	}

	sprintStandupResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Standup.GetBySprint(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.SvcStandup.Key, p, ctx.Logger)), nil
	}

	sprintRetroResolver := func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.GetBySprint(p.Source.(*sprint.Session).ID, paramSetFromGraphQLParams(util.SvcRetro.Key, p, ctx.Logger)), nil
	}

	sprintType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.SvcSprint.Title,
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
				util.KeyOwner: &graphql.Field{
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
				util.Plural(util.KeyMember): &graphql.Field{
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
				util.Plural(util.KeyComment): &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(commentType)),
					Description: "This sprint's comments",
					Args:        listArgs,
					Resolve:     ctxF(sprintCommentResolver),
				},
				util.KeyHistory: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(historyType)),
					Description: "This sprint's name history",
					Args:        listArgs,
					Resolve:     ctxF(sprintHistoryResolver),
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
