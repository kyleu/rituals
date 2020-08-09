package app

import (
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnservice/user"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/socket"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"logur.dev/logur"
)

type Service struct {
	debug    bool
	files    *npncore.FileLoader
	user     *user.Service
	auth     *auth.Service
	Comment  *comment.Service
	Action   *action.Service
	Team     *team.Service
	Sprint   *sprint.Service
	Estimate *estimate.Service
	Standup  *standup.Service
	Retro    *retro.Service
	Socket   *socket.Service
	Database *npndatabase.Service
	version  string
	commit   string
	logger   logur.Logger
}

func NewService(debug bool, db *npndatabase.Service, authEnabled bool, redir string, version string, commitHash string, logger logur.Logger) *Service {
	files := npncore.NewFileLoader("./."+npncore.AppName, logger)
	actionService := action.NewService(db, logger)
	commentService := comment.NewService(actionService, db, logger)
	userSvc := user.NewService(files, db, logger)
	authSvc := auth.NewService(authEnabled, redir, db, logger, userSvc)
	teamSvc := team.NewService(actionService, userSvc, commentService, db, logger)
	sprintSvc := sprint.NewService(actionService, userSvc, commentService, db, logger)
	estimateSvc := estimate.NewService(actionService, userSvc, commentService, db, logger)
	standupSvc := standup.NewService(actionService, userSvc, commentService, db, logger)
	retroSvc := retro.NewService(actionService, userSvc, commentService, db, logger)
	socketSvc := socket.NewService(logger, actionService, userSvc, commentService, authSvc, teamSvc, sprintSvc, estimateSvc, standupSvc, retroSvc)

	return &Service{
		debug:    debug,
		files:    files,
		version:  version,
		commit:   commitHash,
		logger:   logger,
		user:     userSvc,
		auth:     authSvc,
		Comment:  commentService,
		Action:   actionService,
		Team:     teamSvc,
		Sprint:   sprintSvc,
		Estimate: estimateSvc,
		Standup:  standupSvc,
		Retro:    retroSvc,
		Socket:   socketSvc,
		Database: db,
	}
}

func (c *Service) Debug() bool {
	return c.debug
}

func (c *Service) Files() *npncore.FileLoader {
	return c.files
}

func (c *Service) User() *user.Service {
	return c.user
}

func (c *Service) Auth() *auth.Service {
	return c.auth
}

func Comment(a npnweb.AppInfo) *comment.Service {
	return a.(*Service).Comment
}

func Action(a npnweb.AppInfo) *action.Service {
	return a.(*Service).Action
}

func Team(a npnweb.AppInfo) *team.Service {
	return a.(*Service).Team
}

func Sprint(a npnweb.AppInfo) *sprint.Service {
	return a.(*Service).Sprint
}

func Estimate(a npnweb.AppInfo) *estimate.Service {
	return a.(*Service).Estimate
}

func Standup(a npnweb.AppInfo) *standup.Service {
	return a.(*Service).Standup
}

func Retro(a npnweb.AppInfo) *retro.Service {
	return a.(*Service).Retro
}

func Socket(a npnweb.AppInfo) *socket.Service {
	return a.(*Service).Socket
}

func Database(a npnweb.AppInfo) *npndatabase.Service {
	return a.(*Service).Database
}

func (c *Service) Version() string {
	return c.version
}

func (c *Service) Commit() string {
	return c.commit
}

func (c *Service) Logger() logur.Logger {
	return c.logger
}

func (c *Service) Valid() bool {
	return true
}
