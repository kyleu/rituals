package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npngraphql"
)

const MutationName = "RitualsMutation"

func MutationFields() graphql.Fields {
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
			Args:        npngraphql.KeyArgs,
			Resolve:     npngraphql.CtxF(callSandboxResolver),
		},
	}
}
