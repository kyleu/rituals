package web

import (
	"os"
	"strings"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/sessions"
)

type Breadcrumb struct {
	Path  string
	Title string
}

type Breadcrumbs []Breadcrumb

func BreadcrumbsSimple(path string, title string) Breadcrumbs {
	return []Breadcrumb{
		{path, title},
	}
}

type SlugAndTitle interface {
	GetSlug() string
	GetTitle() string
}

var sessionKey = func() string {
	x := os.Getenv("SESSION_KEY")
	if x == "" {
		x = "random_secret_key"
	}
	return x
}()

var store = sessions.NewCookieStore([]byte(sessionKey))

const sessionName = util.AppName + "-session"

func ParseFlash(s string) (string, string) {
	split := strings.SplitN(s, ":", 2)
	severity := split[0]
	content := split[1]

	switch severity {
	case "status":
		return "uk-alert-primary", content
	case "success":
		return "uk-alert-success", content
	case "warning":
		return "uk-alert-warning", content
	case "error":
		return "uk-alert-danger", content
	default:
		return "", content
	}
}
