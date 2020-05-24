package util

import "strings"

const (
	KeyAbout      = "about"
	KeyAction     = "action"
	KeyAdmin      = "admin"
	KeyAuth       = "auth"
	KeyCreated    = "created"
	KeyConnection = "connection"
	KeyDetail     = "detail"
	KeyFeedback   = "feedback"
	KeyGraphQL    = "graphql"
	KeyGraphiQL   = "graphiql"
	KeyID         = "id"
	KeyInvitation = "invitation"
	KeyKey        = "key"
	KeyMember     = "member"
	KeyModules    = "modules"
	KeyNoText     = "-no text-"
	KeyProfile    = "profile"
	KeyReport     = "report"
	KeyRoutes     = "routes"
	KeyPermission = "permission"
	KeySandbox    = "sandbox"
	KeySession    = "session"
	KeySocket     = "socket"
	KeyStatus     = "status"
	KeyStory      = "story"
	KeySvc        = "svc"
	KeyUser       = "user"
	KeyVote       = "vote"
	KeyVoyager    = "voyager"
)

func KeyPlural(k string) string {
	if len(k) == 0 {
		return k
	}
	switch k {
	case KeyStory:
		return "stories"
	case KeySandbox:
		return "sandboxes"
	default:
		return k + "s"
	}
}

func KeyTitle(k string) string {
	if len(k) == 0 {
		return k
	}
	return strings.ToUpper(k[0:1]) + k[1:]
}

func KeyPluralTitle(k string) string {
	return KeyTitle(KeyPlural(k))
}
