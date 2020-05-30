package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/model/action"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	actionResolver     Callback
	actionsResolver    Callback
	actionUserResolver Callback
	actionType         *graphql.Object
)

func initAction() {
	actionResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		id := util.MapGetUUID(p.Args, util.KeyID, ctx.Logger)
		return ctx.App.Action.GetByID(*id), nil
	}

	actionsResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Action.List(paramSetFromGraphQLParams(util.KeyAction, params, ctx.Logger)), nil
	}

	actionUserResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.User.GetByID(p.Source.(*action.Action).UserID, false), nil
	}

	actionType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.Title(util.KeyAction),
			Fields: graphql.Fields{
				util.KeyID: &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				util.KeySvc: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.WithID(util.KeyModel): &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				util.WithID(util.KeyUser): &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				util.KeyAct: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyContent: &graphql.Field{
					Type: graphql.String,
				},
				util.KeyNote: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)

	teamType.AddFieldConfig(util.Plural(util.KeyAction), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))),
		Description: "This sprint's actions",
		Args:        listArgs,
		Resolve:     ctxF(teamActionResolver),
	})

	sprintType.AddFieldConfig(util.Plural(util.KeyAction), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))),
		Description: "This sprint's actions",
		Args:        listArgs,
		Resolve:     ctxF(sprintActionResolver),
	})

	estimateType.AddFieldConfig(util.Plural(util.KeyAction), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))),
		Description: "This estimate's actions",
		Args:        listArgs,
		Resolve:     ctxF(estimateActionResolver),
	})

	standupType.AddFieldConfig(util.Plural(util.KeyAction), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))),
		Description: "This standup's actions",
		Args:        listArgs,
		Resolve:     ctxF(standupActionResolver),
	})

	retroType.AddFieldConfig(util.Plural(util.KeyAction), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))),
		Description: "This retro's actions",
		Args:        listArgs,
		Resolve:     ctxF(retroActionResolver),
	})
}
