package web

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gofrs/uuid"

	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"logur.dev/logur"
)

type RequestContext struct {
	App         *config.AppInfo
	Logger      logur.Logger
	Profile     *util.UserProfile
	Routes      *mux.Router
	Request     *url.URL
	Title       string
	Breadcrumbs Breadcrumbs
	Flashes     []string
	Session     *sessions.Session
}

func (r *RequestContext) Route(act string, pairs ...string) string {
	route := r.Routes.Get(act)
	if route == nil {
		r.App.Logger.Warn("cannot find route at path [" + act + "]")
		return "/routenotfound"
	}
	u, err := route.URL(pairs...)
	if err != nil {
		r.App.Logger.Warn("cannot bind route at path [" + act + "]")
		return "/routeerror"
	}
	return u.Path
}

func ExtractContext(w http.ResponseWriter, r *http.Request, addIfMissing bool) *RequestContext {
	ai, ok := r.Context().Value(util.InfoKey).(*config.AppInfo)
	if !ok {
		ai.Logger.Warn("cannot load AppInfo")
	}
	routes, ok := r.Context().Value(util.RoutesKey).(*mux.Router)
	if !ok {
		ai.Logger.Warn("cannot load Router")
	}
	session, err := store.Get(r, sessionName)
	if err != nil {
		session = sessions.NewSession(store, sessionName)
	}

	var userID uuid.UUID
	userIDValue, ok := session.Values[util.KeyUser]
	if ok && len(userIDValue.(string)) == 36 {
		userID, err = uuid.FromString(userIDValue.(string))
		if err != nil {
			ai.Logger.Warn(fmt.Sprintf("cannot parse uuid [%v]: %+v", userIDValue, err))
			userID = SetSessionUser(util.UUID(), session, r, w, ai.Logger)
		}
	} else {
		userID = SetSessionUser(util.UUID(), session, r, w, ai.Logger)
	}

	user := ai.User.GetByID(userID, addIfMissing)
	var prof *util.UserProfile
	if user == nil {
		prof = util.NewUserProfile(userID)
	} else {
		prof = user.ToProfile()
	}

	flashes := make([]string, 0)
	for _, f := range session.Flashes() {
		flashes = append(flashes, fmt.Sprint(f))
	}

	logger := logur.WithFields(ai.Logger, map[string]interface{}{"path": r.URL.Path, "method": r.Method})

	return &RequestContext{
		App:         ai,
		Logger:      logger,
		Profile:     prof,
		Routes:      routes,
		Request:     r.URL,
		Title:       util.AppName,
		Breadcrumbs: nil,
		Flashes:     flashes,
		Session:     session,
	}
}

func SetSessionUser(userID uuid.UUID, session *sessions.Session, r *http.Request, w http.ResponseWriter, logger logur.Logger) uuid.UUID {
	session.Values[util.KeyUser] = userID.String()
	session.Options = &sessions.Options{Path: "/", HttpOnly: true, SameSite: http.SameSiteStrictMode}
	err := session.Save(r, w)
	if err != nil {
		logger.Warn(fmt.Sprintf("cannot save session: %+v", err))
	}
	return userID
}
