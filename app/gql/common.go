package gql

import (
	"github.com/graphql-go/graphql"
)

var orderingInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Ordering",
		Fields: graphql.InputObjectConfigFieldMap{
			"col": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"asc": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
	},
)

var listArgs = graphql.FieldConfigArgument{
	"orders": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.NewNonNull(orderingInputType)),
	},
	"limit": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"offset": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}
