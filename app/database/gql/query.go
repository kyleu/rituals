package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/util"
)

const QueryName = "RitualsQuery"

func queryFields() graphql.Fields {
	return graphql.Fields{
		util.KeyProfile: &graphql.Field{Type: graphql.NewNonNull(profileType), Resolve: ctxF(profileResolver)},

		util.SvcTeam.Key:    &graphql.Field{Type: teamType, Description: "Get team", Args: keyArgs, Resolve: ctxF(teamResolver)},
		util.SvcTeam.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(teamType))), Description: "Get available teams", Args: listArgs, Resolve: ctxF(teamsResolver)},

		util.SvcSprint.Key:    &graphql.Field{Type: sprintType, Description: "Get sprint", Args: keyArgs, Resolve: ctxF(sprintResolver)},
		util.SvcSprint.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(sprintType))), Description: "Get available sprints", Args: listArgs, Resolve: ctxF(sprintsResolver)},

		util.SvcEstimate.Key:    &graphql.Field{Type: estimateType, Description: "Get estimate", Args: keyArgs, Resolve: ctxF(estimateResolver)},
		util.SvcEstimate.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(estimateType))), Description: "Get available estimates", Args: listArgs, Resolve: ctxF(estimatesResolver)},

		util.KeyStory:              &graphql.Field{Type: storyType, Description: "Get story", Args: idArgs, Resolve: ctxF(storyResolver)},
		util.Plural(util.KeyStory): &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(storyType))), Description: "Get available stories", Args: listArgs, Resolve: ctxF(storiesResolver)},

		util.SvcStandup.Key:    &graphql.Field{Type: standupType, Description: "Get standup", Args: keyArgs, Resolve: ctxF(standupResolver)},
		util.SvcStandup.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(standupType))), Description: "Get available standups", Args: listArgs, Resolve: ctxF(standupsResolver)},

		util.SvcRetro.Key:    &graphql.Field{Type: retroType, Description: "Get retro", Args: keyArgs, Resolve: ctxF(retroResolver)},
		util.SvcRetro.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(retroType))), Description: "Get available retros", Args: listArgs, Resolve: ctxF(retrosResolver)},

		util.KeyUser:              &graphql.Field{Type: profileType, Description: "Get user", Args: idArgs, Resolve: ctxF(userResolver)},
		util.Plural(util.KeyUser): &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(profileType))), Description: "Get available users", Args: listArgs, Resolve: ctxF(usersResolver)},

		util.KeyAction:              &graphql.Field{Type: actionType, Description: "Get action by id", Args: idArgs, Resolve: ctxF(actionResolver)},
		util.Plural(util.KeyAction): &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))), Description: "Get action list", Args: listArgs, Resolve: ctxF(actionsResolver)},

		util.KeySandbox:              &graphql.Field{Type: sandboxType, Description: "Get sandbox by key", Args: keyArgs, Resolve: ctxF(sandboxResolver)},
		util.Plural(util.KeySandbox): &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(sandboxType))), Description: "Get sandbox list", Args: listArgs, Resolve: ctxF(sandboxesResolver)},
	}
}
