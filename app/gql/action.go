package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/action"
)

var (
	actionResolver     Callback
	actionsResolver    Callback
	actionUserResolver Callback
	actionType         *graphql.Object
)

func initAction() {
	actionResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		id := npncore.MapGetUUID(p.Args, npncore.KeyID, ctx.Logger)
		return app.Action(ctx.App).GetByID(*id), nil
	}

	actionsResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Action(ctx.App).List(paramSetFromGraphQLParams(npncore.KeyAction, params, ctx.Logger)), nil
	}

	actionUserResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return ctx.App.User().GetByID(p.Source.(*action.Action).UserID, false), nil
	}

	actionType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: npncore.Title(npncore.KeyAction),
			Fields: graphql.Fields{
				npncore.KeyID: &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				npncore.KeySvc: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.WithID(npncore.KeyModel): &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				npncore.WithID(npncore.KeyUser): &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				npncore.KeyAct: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyContent: &graphql.Field{
					Type: graphql.String,
				},
				npncore.KeyNote: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)

	teamType.AddFieldConfig(npncore.Plural(npncore.KeyAction), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))),
		Description: "This sprint's actions",
		Args:        listArgs,
		Resolve:     ctxF(teamActionResolver),
	})

	sprintType.AddFieldConfig(npncore.Plural(npncore.KeyAction), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))),
		Description: "This sprint's actions",
		Args:        listArgs,
		Resolve:     ctxF(sprintActionResolver),
	})

	estimateType.AddFieldConfig(npncore.Plural(npncore.KeyAction), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))),
		Description: "This estimate's actions",
		Args:        listArgs,
		Resolve:     ctxF(estimateActionResolver),
	})

	standupType.AddFieldConfig(npncore.Plural(npncore.KeyAction), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))),
		Description: "This standup's actions",
		Args:        listArgs,
		Resolve:     ctxF(standupActionResolver),
	})

	retroType.AddFieldConfig(npncore.Plural(npncore.KeyAction), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))),
		Description: "This retro's actions",
		Args:        listArgs,
		Resolve:     ctxF(retroActionResolver),
	})
}
