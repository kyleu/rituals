package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	historyType *graphql.Object
)

func initHistory() {
	historyType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.Title(util.KeyHistory),
			Fields: graphql.Fields{
				util.KeySlug: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.WithID(util.KeyModel): &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				"modelName": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)
}
