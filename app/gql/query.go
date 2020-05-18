package gql

import "github.com/graphql-go/graphql"

func queryFields() graphql.Fields {
	return graphql.Fields{
		"profile": &graphql.Field{Type: graphql.NewNonNull(profileType), Resolve: ctxF(profileResolver)},

		"estimate":  &graphql.Field{Type: estimateType, Description: "Get estimate", Args: estimateArgs, Resolve: ctxF(estimateResolver)},
		"estimates": &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(estimateType))), Description: "Get available estimates", Args: listArgs, Resolve: ctxF(estimatesResolver)},

		"standup":  &graphql.Field{Type: standupType, Description: "Get standup", Args: standupArgs, Resolve: ctxF(standupResolver)},
		"standups": &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(standupType))), Description: "Get available standups", Args: listArgs, Resolve: ctxF(standupsResolver)},

		"retro":  &graphql.Field{Type: retroType, Description: "Get retro", Args: retroArgs, Resolve: ctxF(retroResolver)},
		"retros": &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(retroType))), Description: "Get available retros", Args: listArgs, Resolve: ctxF(retrosResolver)},

		"sprint":  &graphql.Field{Type: sprintType, Description: "Get sprint", Args: sprintArgs, Resolve: ctxF(sprintResolver)},
		"sprints": &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(sprintType))), Description: "Get available sprints", Args: listArgs, Resolve: ctxF(sprintsResolver)},

		"sandbox":   &graphql.Field{Type: sandboxType, Description: "Get sandbox by key", Args: sandboxArgs, Resolve: ctxF(sandboxResolver)},
		"sandboxes": &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(sandboxType))), Description: "Get sandbox list", Args: listArgs, Resolve: ctxF(sandboxesResolver)},
	}
}
