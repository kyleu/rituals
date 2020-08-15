package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app/sandbox"
)

var (
	sandboxResolver     npngraphql.Callback
	sandboxesResolver   npngraphql.Callback
	callSandboxResolver npngraphql.Callback
	sandboxType         *graphql.Object
)

func initSandbox() {
	sandboxResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return sandbox.FromString(npncore.MapGetString(p.Args, npncore.KeyKey, ctx.Logger)), nil
	}

	sandboxesResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return sandbox.AllSandboxes, nil
	}

	callSandboxResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		key := npncore.MapGetString(p.Args, npncore.KeyKey, ctx.Logger)
		sb := sandbox.FromString(key)
		if sb == nil {
			return nil, npncore.IDError(npncore.KeySandbox, key)
		}
		_, rsp, err := sb.Resolve(ctx)
		return rsp, err
	}

	sandboxType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: npncore.Title(npncore.KeySandbox),
			Fields: graphql.Fields{
				npncore.KeyKey: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyTitle: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
