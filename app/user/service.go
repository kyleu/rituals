package user

import (
	"database/sql"
	"time"

	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	actions *action.Service
	db      *sqlx.DB
	logger  logur.Logger
}

func NewService(actions *action.Service, db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "user"})

	return &Service{
		actions: actions,
		db:      db,
		logger:  logger,
	}
}

func (s *Service) New(id uuid.UUID) (*SystemUser, error) {
	s.logger.Info("creating user [" + id.String() + "]")
	q := "insert into system_user (id, name, role, theme, nav_color, link_color, picture, locale, created) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	prof := util.NewUserProfile(id)
	_, err := s.db.Exec(q, prof.UserID, prof.Name, util.RoleGuest.Key, prof.Theme.String(), prof.NavColor, prof.LinkColor, prof.Picture, prof.Locale.String(), time.Now())
	if err != nil {
		return nil, err
	}
	return s.GetByID(id, false)
}

func (s *Service) List(params *query.Params) ([]*SystemUser, error) {
	params = query.ParamsWithDefaultOrdering("user", params, &query.Ordering{Column: "created", Asc: false})
	var ret []*SystemUser
	err := s.db.Select(&ret, query.SQLSelect("*", "system_user", "", params.OrderByString(), params.Limit, params.Offset))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetByID(id uuid.UUID, addIfMissing bool) (*SystemUser, error) {
	ret := &SystemUser{}
	err := s.db.Get(ret, query.SQLSelect("*", "system_user", "id = $1", "", 0, 0), id)
	if err == sql.ErrNoRows {
		if addIfMissing {
			return s.New(id)
		} else {
			return nil, nil
		}
	}
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *Service) SaveProfile(prof *util.UserProfile) (*util.UserProfile, error) {
	s.logger.Info("updating user [" + prof.UserID.String() + "] from profile")
	q := "update system_user set name = $1, role = $2, theme = $3, nav_color = $4, link_color = $5, picture = $6, locale = $7 where id = $8"
	_, err := s.db.Exec(q, prof.Name, prof.Role.Key, prof.Theme.String(), prof.NavColor, prof.LinkColor, prof.Picture, prof.Locale.String(), prof.UserID)
	if err != nil {
		return nil, err
	}
	return prof, nil
}

func (s *Service) UpdateUserName(id uuid.UUID, name string) error {
	s.logger.Info("updating user [" + id.String() + "]")
	q := "update system_user set name = $1 where id = $2"
	_, err := s.db.Exec(q, name, id)
	return err
}
