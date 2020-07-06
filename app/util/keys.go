package util

import "strings"

const (
	KeyAbout      = "about"
	KeyAct        = "act"
	KeyAction     = "action"
	KeyAdmin      = "admin"
	KeyAuth       = "auth"
	KeyCreated    = "created"
	KeyComment    = "comment"
	KeyConnection = "connection"
	KeyCategory   = "category"
	KeyChoice     = "choice"
	KeyContent    = "content"
	KeyDetail     = "detail"
	KeyEmail      = "email"
	KeyError      = "error"
	KeyExport     = "export"
	KeyFeedback   = "feedback"
	KeyFmt        = "fmt"
	KeyGraphQL    = "graphql"
	KeyGraphiQL   = "graphiql"
	KeyHistory    = "history"
	KeyHTML       = "html"
	KeyID         = "id"
	KeyIdx        = "idx"
	KeyInvitation = "invitation"
	KeyJSON       = "json"
	KeyKey        = "key"
	KeyMember     = "member"
	KeyMigration  = "migration"
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
	KeyTranscript = "transcript"
	KeyUser       = "user"
	KeyVote       = "vote"
	KeyVoyager    = "voyager"
)

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
	return Plural(Title(k))
}

func WithID(k string) string {
	return k + "ID"
}

func WithDBID(k string) string {
	return k + "_id"
}
