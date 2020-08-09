package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
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

var idArgs = graphql.FieldConfigArgument{
	npncore.KeyID: &graphql.ArgumentConfig{
		Type: scalarUUID,
	},
}

var keyArgs = graphql.FieldConfigArgument{
	npncore.KeyKey: &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
