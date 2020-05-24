package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/sandbox"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	sandboxArgs         graphql.FieldConfigArgument
	sandboxResolver     Callback
	sandboxesResolver   Callback
	callSandboxArgs     graphql.FieldConfigArgument
	callSandboxResolver Callback
	sandboxType         *graphql.Object
)

func initSandbox() {
	sandboxArgs = graphql.FieldConfigArgument{
		util.KeyKey: &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}

	sandboxResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		key, err := paramString(p, util.KeyKey)
		if err != nil {
			return nil, err
		}
		return sandbox.FromString(key), nil
	}

	sandboxesResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return sandbox.AllSandboxes, nil
	}

	callSandboxArgs = graphql.FieldConfigArgument{
		util.KeyKey: &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}

	callSandboxResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		key, _ := paramString(params, util.KeyKey)
		sb := sandbox.FromString(key)
		if sb == nil {
			return nil, util.IDError(util.KeySandbox, key)
		}
		return sb.Resolve(ctx)
	}

	sandboxType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Sandbox",
			Fields: graphql.Fields{
				util.KeyKey: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"title": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
