package util

import "strings"

const (
	KeyAbout      = "about"
	KeyAct        = "act"
	KeyAction     = "action"
	KeyAdmin      = "admin"
	KeyAuth       = "auth"
	KeyAuthor     = "author"
	KeyCreated    = "created"
	KeyComment    = "comment"
	KeyConnection = "connection"
	KeyCategory   = "category"
	KeyChoice     = "choice"
	KeyContent    = "content"
	KeyDetail     = "detail"
	KeyEmail      = "email"
	KeyError      = "error"
	KeyFeedback   = "feedback"
	KeyGraphQL    = "graphql"
	KeyGraphiQL   = "graphiql"
	KeyHTML       = "html"
	KeyID         = "id"
	KeyIdx        = "idx"
	KeyInvitation = "invitation"
	KeyKey        = "key"
	KeyMember     = "member"
	KeyModel      = "model"
	KeyModules    = "modules"
	KeyName       = "name"
	KeyNote       = "note"
	KeyNoText     = "-no text-"
	KeyProfile    = "profile"
	KeyOwner      = "owner"
	KeyReport     = "report"
	KeyRole       = "role"
	KeyRoutes     = "routes"
	KeyPermission = "permission"
	KeyProvider   = "provider"
	KeySandbox    = "sandbox"
	KeyService    = "service"
	KeySession    = "session"
	KeySlug       = "slug"
	KeySocket     = "socket"
	KeyStatus     = "status"
	KeyStory      = "story"
	KeySvc        = "svc"
	KeySystem     = "system"
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
	case KeyCategory:
		return "categories"
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

func PluralTitle(k string) string {
	return Title(Plural(k))
}

func WithID(k string) string {
	return k + "ID"
}

func WithDBID(k string) string {
	return k + "_id"
}
