package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/sandbox"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	sandboxResolver     Callback
	sandboxesResolver   Callback
	callSandboxResolver Callback
	sandboxType         *graphql.Object
)

func initSandbox() {
	sandboxResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return sandbox.FromString(util.MapGetString(p.Args, util.KeyKey, ctx.Logger)), nil
	}

	sandboxesResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return sandbox.AllSandboxes, nil
	}

	callSandboxResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		key := util.MapGetString(p.Args, util.KeyKey, ctx.Logger)
		sb := sandbox.FromString(key)
		if sb == nil {
			return nil, util.IDError(util.KeySandbox, key)
		}
		return sb.Resolve(ctx)
	}

	sandboxType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.Title(util.KeySandbox),
			Fields: graphql.Fields{
				util.KeyKey: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyTitle: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
