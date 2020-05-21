package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/util"
)

const QueryName = "RitualsQuery"

func queryFields() graphql.Fields {
	return graphql.Fields{
		util.KeyProfile: &graphql.Field{Type: graphql.NewNonNull(profileType), Resolve: ctxF(profileResolver)},

		util.SvcTeam.Key:    &graphql.Field{Type: teamType, Description: "Get team", Args: teamArgs, Resolve: ctxF(teamResolver)},
		util.SvcTeam.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(teamType))), Description: "Get available teams", Args: listArgs, Resolve: ctxF(teamsResolver)},

		util.SvcSprint.Key:    &graphql.Field{Type: sprintType, Description: "Get sprint", Args: sprintArgs, Resolve: ctxF(sprintResolver)},
		util.SvcSprint.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(sprintType))), Description: "Get available sprints", Args: listArgs, Resolve: ctxF(sprintsResolver)},

		util.SvcEstimate.Key:    &graphql.Field{Type: estimateType, Description: "Get estimate", Args: estimateArgs, Resolve: ctxF(estimateResolver)},
		util.SvcEstimate.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(estimateType))), Description: "Get available estimates", Args: listArgs, Resolve: ctxF(estimatesResolver)},

		util.SvcStandup.Key:    &graphql.Field{Type: standupType, Description: "Get standup", Args: standupArgs, Resolve: ctxF(standupResolver)},
		util.SvcStandup.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(standupType))), Description: "Get available standups", Args: listArgs, Resolve: ctxF(standupsResolver)},

		util.SvcRetro.Key:    &graphql.Field{Type: retroType, Description: "Get retro", Args: retroArgs, Resolve: ctxF(retroResolver)},
		util.SvcRetro.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(retroType))), Description: "Get available retros", Args: listArgs, Resolve: ctxF(retrosResolver)},

		util.KeyUser: &graphql.Field{Type: profileType, Description: "Get user", Args: userArgs, Resolve: ctxF(userResolver)},
		"users":      &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(profileType))), Description: "Get available users", Args: listArgs, Resolve: ctxF(usersResolver)},

		util.KeySandbox: &graphql.Field{Type: sandboxType, Description: "Get sandbox by key", Args: sandboxArgs, Resolve: ctxF(sandboxResolver)},
		"sandboxes":     &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(sandboxType))), Description: "Get sandbox list", Args: listArgs, Resolve: ctxF(sandboxesResolver)},
	}
}
