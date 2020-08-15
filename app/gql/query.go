package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/rituals.dev/app/util"
)

const QueryName = "RitualsQuery"

func QueryFields() graphql.Fields {
	return graphql.Fields{
		npncore.KeyProfile: &graphql.Field{Type: graphql.NewNonNull(profileType), Resolve: npngraphql.CtxF(profileResolver)},

		util.SvcTeam.Key:    &graphql.Field{Type: teamType, Description: "Get team", Args: npngraphql.KeyArgs, Resolve: npngraphql.CtxF(teamResolver)},
		util.SvcTeam.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(teamType))), Description: "Get available teams", Args: npngraphql.ListArgs, Resolve: npngraphql.CtxF(teamsResolver)},

		util.SvcSprint.Key:    &graphql.Field{Type: sprintType, Description: "Get sprint", Args: npngraphql.KeyArgs, Resolve: npngraphql.CtxF(sprintResolver)},
		util.SvcSprint.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(sprintType))), Description: "Get available sprints", Args: npngraphql.ListArgs, Resolve: npngraphql.CtxF(sprintsResolver)},

		util.SvcEstimate.Key:    &graphql.Field{Type: estimateType, Description: "Get estimate", Args: npngraphql.KeyArgs, Resolve: npngraphql.CtxF(estimateResolver)},
		util.SvcEstimate.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(estimateType))), Description: "Get available estimates", Args: npngraphql.ListArgs, Resolve: npngraphql.CtxF(estimatesResolver)},

		util.KeyStory:                 &graphql.Field{Type: storyType, Description: "Get story", Args: npngraphql.IDArgs, Resolve: npngraphql.CtxF(storyResolver)},
		npncore.Plural(util.KeyStory): &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(storyType))), Description: "Get available stories", Args: npngraphql.ListArgs, Resolve: npngraphql.CtxF(storiesResolver)},

		util.SvcStandup.Key:    &graphql.Field{Type: standupType, Description: "Get standup", Args: npngraphql.KeyArgs, Resolve: npngraphql.CtxF(standupResolver)},
		util.SvcStandup.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(standupType))), Description: "Get available standups", Args: npngraphql.ListArgs, Resolve: npngraphql.CtxF(standupsResolver)},

		util.SvcRetro.Key:    &graphql.Field{Type: retroType, Description: "Get retro", Args: npngraphql.KeyArgs, Resolve: npngraphql.CtxF(retroResolver)},
		util.SvcRetro.Plural: &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(retroType))), Description: "Get available retros", Args: npngraphql.ListArgs, Resolve: npngraphql.CtxF(retrosResolver)},

		npncore.KeyUser:                 &graphql.Field{Type: profileType, Description: "Get user", Args: npngraphql.IDArgs, Resolve: npngraphql.CtxF(userResolver)},
		npncore.Plural(npncore.KeyUser): &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(profileType))), Description: "Get available users", Args: npngraphql.ListArgs, Resolve: npngraphql.CtxF(usersResolver)},

		npncore.KeyAction:                 &graphql.Field{Type: actionType, Description: "Get action by id", Args: npngraphql.IDArgs, Resolve: npngraphql.CtxF(actionResolver)},
		npncore.Plural(npncore.KeyAction): &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(actionType))), Description: "Get action list", Args: npngraphql.ListArgs, Resolve: npngraphql.CtxF(actionsResolver)},

		npncore.KeySandbox:                 &graphql.Field{Type: sandboxType, Description: "Get sandbox by key", Args: npngraphql.KeyArgs, Resolve: npngraphql.CtxF(sandboxResolver)},
		npncore.Plural(npncore.KeySandbox): &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(sandboxType))), Description: "Get sandbox list", Args: npngraphql.ListArgs, Resolve: npngraphql.CtxF(sandboxesResolver)},
	}
}
