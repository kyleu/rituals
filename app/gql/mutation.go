package gql

import (
	"github.com/graphql-go/graphql"
)

var mutationFields = graphql.Fields{
	"ping": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "pong", nil
		},
	},
}
