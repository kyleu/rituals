package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/web"
)

var graphQLInitialized = false

type Callback func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error)

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

func paramSetFromGraphQLParams(key string, params graphql.ResolveParams) *query.Params {
	orderings := make([]*query.Ordering, 0)
	o, ok := params.Args["orders"]
	if ok {
		for _, x := range o.([]interface{}) {
			m := x.(map[string]interface{})
			col := m["col"].(string)
			orderings = append(orderings, &query.Ordering{Column: col, Asc: m["asc"].(bool)})
		}
	}

	limit := 0
	l, ok := params.Args["limit"]
	if ok {
		limit = l.(int)
	}

	offset := 0
	x, ok := params.Args["offset"]
	if ok {
		offset = x.(int)
	}

	ret := &query.Params{Key: key, Orderings: orderings, Limit: limit, Offset: offset}
	return ret.Filtered()
}
