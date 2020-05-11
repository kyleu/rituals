package user

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	db     *sqlx.DB
	logger logur.Logger
}

func NewService(db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "user"})

	return &Service{
		db:     db,
		logger: logger,
	}
}

func (s *Service) List() ([]SystemUser, error) {
	var ret []SystemUser
	err := s.db.Select(&ret, "select * from system_user")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetByID(id uuid.UUID, addIfMissing bool) (*SystemUser, error) {
	ret := &SystemUser{}
	err := s.db.Get(ret, "select * from system_user where id = $1", id)
	if err == sql.ErrNoRows {
		if addIfMissing {
			return s.CreateNewUser(id)
		} else {
			return nil, nil
		}
	}
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *Service) CreateNewUser(id uuid.UUID) (*SystemUser, error) {
	s.logger.Info("creating user [" + id.String() + "]")
	q := "insert into system_user (id, name, role, theme, nav_color, link_color, locale, created) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	role := "guest"
	prof := util.NewUserProfile(id)
	_, err := s.db.Exec(q, prof.UserID, prof.Name, role, prof.Theme.String(), prof.NavColor, prof.LinkColor, prof.Locale.String(), time.Now())
	if err != nil {
		return nil, err
	}
	return s.GetByID(id, false)
}

func (s *Service) SaveProfile(prof *util.UserProfile) (*util.UserProfile, error) {
	s.logger.Info("updating user [" + prof.UserID.String() + "] from profile")
	q := "update system_user set name = $1, role = $2, theme = $3, nav_color = $4, link_color = $5, locale = $6 where id = $7"
	role := "guest"
	_, err := s.db.Exec(q, prof.Name, role, prof.Theme.String(), prof.NavColor, prof.LinkColor, prof.Locale.String(), prof.UserID)
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
