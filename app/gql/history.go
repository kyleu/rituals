package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
)

var (
	historyType *graphql.Object
)

func initHistory() {
	historyType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: npncore.Title(npncore.KeyHistory),
			Fields: graphql.Fields{
				npncore.KeySlug: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.WithID(npncore.KeyModel): &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				"modelName": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)
}
