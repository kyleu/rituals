package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/util"
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
	util.KeyID: &graphql.ArgumentConfig{
		Type: scalarUUID,
	},
}

var keyArgs = graphql.FieldConfigArgument{
	util.KeyKey: &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
