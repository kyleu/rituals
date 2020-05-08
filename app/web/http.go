package web

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"logur.dev/logur"
)

type Breadcrumb struct {
	Path  string
	Title string
}

type Breadcrumbs []Breadcrumb

func (bc Breadcrumbs) Title(ai *config.AppInfo) string {
	if len(bc) == 0 {
		return util.AppName
	}
	return bc[len(bc)-1].Title
}

func BreadcrumbsSimple(path string, title string) Breadcrumbs {
	return []Breadcrumb{
		{path, title},
	}
}

type RequestContext struct {
	App         *config.AppInfo
	Logger      logur.LoggerFacade
	Profile     *util.UserProfile
	Routes      *mux.Router
	Title       string
	Breadcrumbs Breadcrumbs
	Flashes     []string
	Session     sessions.Session
}

func (r *RequestContext) Route(act string, pairs ...string) string {
	route := r.Routes.Get(act)
	if route == nil {
		r.App.Logger.Warn("cannot find route at path [" + act + "]")
		return "/routenotfound"
	}
	url, err := route.URL(pairs...)
	if err != nil {
		r.App.Logger.Warn("cannot bind route at path [" + act + "]")
		return "/routeerror"
	}
	return url.Path
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

func ExtractContext(w http.ResponseWriter, r *http.Request) RequestContext {
	ai := r.Context().Value("info").(*config.AppInfo)
	routes := r.Context().Value("routes").(*mux.Router)
	session, err := store.Get(r, sessionName)
	if err != nil {
		session = sessions.NewSession(store, sessionName)
	}

	var userID uuid.UUID
	userIDValue, ok := session.Values["user"]
	if ok {
		userID, err = uuid.FromString(userIDValue.(string))
		if err != nil {
			ai.Logger.Warn(fmt.Sprintf("cannot parse uuid [%v]: %+v", userIDValue, err))
		}
	} else {
		userID = util.UUID()
		_, err := ai.User.CreateNewUser(userID)
		if err != nil {
			ai.Logger.Warn(fmt.Sprintf("cannot save user: %+v", err))
		}
		session.Values["user"] = userID.String()
		err = session.Save(r, w)
		if err != nil {
			ai.Logger.Warn(fmt.Sprintf("cannot save session: %+v", err))
		}
	}

	user, err := ai.User.GetByID(userID, true)
	if err != nil {
		ai.Logger.Warn(fmt.Sprintf("unable to load user profile: %v", err))
	}
	var prof *util.UserProfile
	if user == nil {
		fallback := util.NewUserProfile(userID)
		prof = &fallback
	} else {
		fallback := user.ToProfile()
		prof = &fallback
	}

	flashes := make([]string, 0)
	for _, f := range session.Flashes() {
		flashes = append(flashes, fmt.Sprintf("%v", f))
	}

	logger := logur.WithFields(ai.Logger, map[string]interface{}{"path": r.URL.Path, "method": r.Method})

	return RequestContext{
		App:         ai,
		Logger:      logger,
		Profile:     prof,
		Routes:      routes,
		Title:       util.AppName,
		Breadcrumbs: nil,
		Flashes:     flashes,
		Session:     *session,
	}
}

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
