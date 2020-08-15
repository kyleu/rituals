package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/rituals.dev/app/permission"
)

var (
	permissionType *graphql.Object
)

func initPermission() {
	permissionType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: npncore.Title(npncore.KeyPermission),
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
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)

	estimateType.AddFieldConfig(npncore.Plural(npncore.KeyPermission), &graphql.Field{
		Type:        graphql.NewList(graphql.NewNonNull(permissionType)),
		Description: "This estimate's permissions",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(estimatePermissionResolver),
	})

	standupType.AddFieldConfig(npncore.Plural(npncore.KeyPermission), &graphql.Field{
		Type:        graphql.NewList(graphql.NewNonNull(permissionType)),
		Description: "This standup's permissions",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(standupPermissionResolver),
	})

	retroType.AddFieldConfig(npncore.Plural(npncore.KeyPermission), &graphql.Field{
		Type:        graphql.NewList(graphql.NewNonNull(permissionType)),
		Description: "This retro's permissions",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(retroPermissionResolver),
	})

	sprintType.AddFieldConfig(npncore.Plural(npncore.KeyPermission), &graphql.Field{
		Type:        graphql.NewList(graphql.NewNonNull(permissionType)),
		Description: "This sprint's permissions",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(sprintPermissionResolver),
	})

	teamType.AddFieldConfig(npncore.Plural(npncore.KeyPermission), &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(permissionType))),
		Description: "This team's permissions",
		Args:        npngraphql.ListArgs,
		Resolve:     npngraphql.CtxF(teamPermissionResolver),
	})
}
