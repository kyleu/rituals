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
	KeyIdx        = "idx"
	KeyInvitation = "invitation"
	KeyKey        = "key"
	KeyMember     = "member"
	KeyModules    = "modules"
	KeyName       = "name"
	KeyNoText     = "-no text-"
	KeyProfile    = "profile"
	KeyOwner      = "owner"
	KeyReport     = "report"
	KeyRole       = "role"
	KeyRoutes     = "routes"
	KeyPermission = "permission"
	KeySandbox    = "sandbox"
	KeyService    = "service"
	KeySession    = "session"
	KeySlug       = "slug"
	KeySocket     = "socket"
	KeyStatus     = "status"
	KeyStory      = "story"
	KeySvc        = "svc"
	KeySystemUser = "system_user"
	KeyTheme      = "theme"
	KeyTitle      = "title"
	KeyUser       = "user"
	KeyVote       = "vote"
	KeyVoyager    = "voyager"
)

func Plural(k string) string {
	if len(k) == 0 {
		return k
	}
	switch k {
	case KeyGraphQL, KeyRoutes, KeyModules:
		return k
	case KeyStory:
		return "stories"
	case KeySandbox:
		return "sandboxes"
	default:
		return k + "s"
	}
}

func Title(k string) string {
	if len(k) == 0 {
		return k
	}
	switch k {
	case KeyID:
		return "ID"
	case KeyIdx:
		return "Index"
	case KeyGraphQL:
		return "GraphQL"
	}
	return strings.ToUpper(k[0:1]) + k[1:]
}

func PluralProper(k string) string {
	return Title(Plural(k))
}

func WithID(k string) string {
	return k + "ID"
}
