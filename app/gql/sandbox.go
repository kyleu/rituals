package gql

import (
	"emperror.dev/errors"
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/sandbox"
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
		"key": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}

	sandboxResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		key, ok := p.Args["key"].(string)
		if ok {
			return sandbox.SandboxFromString(key), nil
		}
		return nil, nil
	}

	sandboxesResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return sandbox.AllSandboxes, nil
	}

	callSandboxArgs = graphql.FieldConfigArgument{
		"key": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}

	callSandboxResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		key, _ := params.Args["key"].(string)
		sb := sandbox.SandboxFromString(key)
		if sb == nil {
			return "", errors.New("invalid sandbox [" + key + "]")
		}
		return sb.Resolve(ctx)
	}

	sandboxType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Sandbox",
			Fields: graphql.Fields{
				"key": &graphql.Field{
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
