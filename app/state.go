package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/auth"
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filesystem"
	"github.com/kyleu/rituals/app/lib/log"
	"github.com/kyleu/rituals/app/lib/telemetry"
	"github.com/kyleu/rituals/app/lib/theme"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
)

var once sync.Once

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func (b *BuildInfo) String() string {
	if b.Date == util.KeyUnknown {
		return b.Version
	}
	d, _ := util.TimeFromJS(b.Date)
	return fmt.Sprintf("%s (%s)", b.Version, util.TimeToYMD(d))
}

type State struct {
	Debug     bool
	BuildInfo *BuildInfo
	Files     filesystem.FileLoader
	Auth      *auth.Service
	DB        *database.Service
	Themes    *theme.Service
	Services  *Services
	Started   time.Time
}

func NewState(ctx context.Context, debug bool, bi *BuildInfo, f filesystem.FileLoader, enableTelemetry bool, port uint16, logger util.Logger) (*State, error) {
	var loadLocationError error
	once.Do(func() {
		loc, err := time.LoadLocation("UTC")
		if err != nil {
			loadLocationError = err
			return
		}
		time.Local = loc
	})
	if loadLocationError != nil {
		return nil, loadLocationError
	}

	_ = telemetry.InitializeIfNeeded(ctx, enableTelemetry, bi.Version, logger)

	return &State{
		Debug:     debug,
		BuildInfo: bi,
		Files:     f,
		Auth:      auth.NewService("", port, logger),
		Themes:    theme.NewService(f),
		Started:   util.TimeCurrent(),
	}, nil
}

func (s *State) Close(ctx context.Context, logger util.Logger) error {
	defer func() { _ = telemetry.Close(ctx) }()
	if err := s.DB.Close(); err != nil {
		logger.Errorf("error closing database: %+v", err)
	}
	return s.Services.Close(ctx, logger)
}

func (s *State) User(ctx context.Context, id uuid.UUID, logger util.Logger) (*user.User, error) {
	if s == nil || s.Services == nil || s.Services.User == nil {
		return nil, nil
	}
	return s.Services.User.Get(ctx, nil, id, logger)
}

func Bootstrap(ctx context.Context, bi *BuildInfo, cfgDir string, port uint16, debug bool, logger util.Logger) (*State, error) {
	fs, err := filesystem.NewFileSystem(cfgDir, false, "")
	if err != nil {
		return nil, err
	}

	telemetryDisabled := util.GetEnvBoolAny(false, "disable_telemetry", "telemetry_disabled")
	st, err := NewState(ctx, debug, bi, fs, !telemetryDisabled, port, logger)
	if err != nil {
		return nil, err
	}

	ctx, span, logger := telemetry.StartSpan(ctx, "app:init", logger)
	defer span.Complete()
	t := util.TimerStart()

	db, err := database.OpenDefaultPostgres(ctx, logger)
	if err != nil {
		logger.Errorf("unable to open default database: %+v", err)
	}
	st.DB = db
	svcs, err := NewServices(ctx, st, logger)
	if err != nil {
		return nil, err
	}
	logger.Debugf("created app state in [%s]", util.MicrosToMillis(t.End()))
	st.Services = svcs

	return st, nil
}

func BootstrapRunDefault[T any](ctx context.Context, bi *BuildInfo, fn func(as *State, logger util.Logger) (T, error)) (T, error) {
	logger, _ := log.InitLogging(false)
	as, err := Bootstrap(ctx, bi, util.ConfigDir, 0, false, logger)
	if err != nil {
		var dflt T
		return dflt, err
	}
	ret, err := fn(as, logger)
	if err != nil {
		var dflt T
		return dflt, err
	}
	err = as.Close(ctx, logger)
	if err != nil {
		var dflt T
		return dflt, err
	}
	return ret, nil
}
