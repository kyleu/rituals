package gql

import (
	"github.com/graphql-go/graphql"
)

func mutationFields() graphql.Fields {
	return graphql.Fields{
		"ping": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "pong", nil
			},
		},

		"callSandbox": &graphql.Field{
			Type:        graphql.String,
			Description: "Call sandbox",
			Args:        callSandboxArgs,
			Resolve:     ctxF(callSandboxResolver),
		},
	}
}
