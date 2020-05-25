package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	permissionType           *graphql.Object
)

func initPermission() {
	permissionType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.KeyTitle(util.KeyPermission),
			Fields: graphql.Fields{
				"k": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"v": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"access": &graphql.Field{
					Type: graphql.NewNonNull(memberRoleType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*permission.Permission).Access.Key, nil
					},
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)

	estimateType.AddFieldConfig(util.KeyPlural(util.KeyPermission), &graphql.Field{
		Type:        graphql.NewList(graphql.NewNonNull(permissionType)),
		Description: "This estimate's permissions",
		Args:        listArgs,
		Resolve:     ctxF(estimatePermissionResolver),
	})

	standupType.AddFieldConfig(util.KeyPlural(util.KeyPermission), &graphql.Field{
		Type:        graphql.NewList(graphql.NewNonNull(permissionType)),
		Description: "This standup's permissions",
		Args:        listArgs,
		Resolve:     ctxF(standupPermissionResolver),
	})

	retroType.AddFieldConfig(util.KeyPlural(util.KeyPermission), &graphql.Field{
		Type:        graphql.NewList(graphql.NewNonNull(permissionType)),
		Description: "This retro's permissions",
		Args:        listArgs,
		Resolve:     ctxF(retroPermissionResolver),
	})

	sprintType.AddFieldConfig(util.KeyPlural(util.KeyPermission), &graphql.Field{
		Type:        graphql.NewList(graphql.NewNonNull(permissionType)),
		Description: "This sprint's permissions",
		Args:        listArgs,
		Resolve:     ctxF(sprintPermissionResolver),
	})

	teamType.AddFieldConfig(util.KeyPlural(util.KeyPermission), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(permissionType))),
		Description: "This team's permissions",
		Args:        listArgs,
		Resolve:     ctxF(teamPermissionResolver),
	})
}
