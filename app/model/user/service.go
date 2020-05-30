package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kyleu/rituals.dev/app/database"

	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/action"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	actions *action.Service
	db      *database.Service
	logger  logur.Logger
}

func NewService(actions *action.Service, db *database.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.KeyUser})

	return &Service{
		actions: actions,
		db:      db,
		logger:  logger,
	}
}

func (s *Service) new(id uuid.UUID) (*SystemUser, error) {
	s.logger.Info("creating user [" + id.String() + "]")

	q := query.SQLInsert(util.KeySystemUser, []string{util.KeyID, util.KeyName, util.KeyRole, util.KeyTheme, "nav_color", "link_color", "picture", "locale", util.KeyCreated}, 1)
	prof := util.NewUserProfile(id)
	err := s.db.Insert(q, nil, prof.UserID, prof.Name, util.RoleGuest.Key, prof.Theme.String(), prof.NavColor, prof.LinkColor, prof.Picture, prof.Locale.String(), time.Now())

	if err != nil {
		return nil, err
	}

	return s.GetByID(id, false), nil
}

func (s *Service) List(params *query.Params) SystemUsers {
	params = query.ParamsWithDefaultOrdering(util.KeyUser, params, query.DefaultCreatedOrdering...)

	var ret SystemUsers

	q := query.SQLSelect("*", util.KeySystemUser, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&ret, q, nil)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving system users: %+v", err))
		return nil
	}

	return ret
}

func (s *Service) GetByID(id uuid.UUID, addIfMissing bool) *SystemUser {
	ret := &SystemUser{}
	q := query.SQLSelectSimple("*", util.KeySystemUser, util.KeyID + " = $1")
	err := s.db.Get(ret, q, nil, id)
	if err == sql.ErrNoRows {
		if addIfMissing {
			ret, err := s.new(id)
			if err != nil {
				util.LogError(s.logger, "error creating new user with id [%v]: %+v", id, err)
			}
			return ret
		}
		return nil
	}
	if err != nil {
		util.LogError(s.logger, "error getting user by id [%v]: %+v", id, err)
		return nil
	}
	return ret
}

func (s *Service) SaveProfile(prof *util.UserProfile) (*util.UserProfile, error) {
	s.logger.Info("updating user [" + prof.UserID.String() + "] from profile")
	cols := []string{"name", "role", "theme", "nav_color", "link_color", "picture", "locale"}
	q := query.SQLUpdate(util.KeySystemUser, cols, fmt.Sprintf("%v = $%v", util.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, prof.Name, prof.Role.Key, prof.Theme.String(), prof.NavColor, prof.LinkColor, prof.Picture, prof.Locale.String(), prof.UserID)
	if err != nil {
		return nil, err
	}
	return prof, nil
}

func (s *Service) UpdateUserName(userID uuid.UUID, name string) error {
	s.logger.Info("updating user [" + userID.String() + "]")
	cols := []string{"name"}
	q := query.SQLUpdate(util.KeySystemUser, cols, fmt.Sprintf("%v = $%v", util.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, name, userID)
	return err
}

func (s *Service) SetRole(userID uuid.UUID, role util.Role) error {
	_ = s.GetByID(userID, true)
	s.logger.Info("updating user role [" + userID.String() + "]")
	cols := []string{"role"}
	q := query.SQLUpdate(util.KeySystemUser, cols, fmt.Sprintf("%v = $%v", util.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, role.Key, userID)
	return err
}
