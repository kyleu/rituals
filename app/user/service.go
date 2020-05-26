package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kyleu/rituals.dev/app/database"

	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"

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

	q := query.SQLInsert(util.KeySystemUser, []string{"id", "name", "role", "theme", "nav_color", "link_color", "picture", "locale", "created"}, 1)
	prof := util.NewUserProfile(id)
	err := s.db.Insert(q, nil, prof.UserID, prof.Name, util.RoleGuest.Key, prof.Theme.String(), prof.NavColor, prof.LinkColor, prof.Picture, prof.Locale.String(), time.Now())

	if err != nil {
		return nil, err
	}

	return s.GetByID(id, false)
}

func (s *Service) List(params *query.Params) (SystemUsers, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyUser, params, query.DefaultCreatedOrdering...)

	var ret SystemUsers

	q := query.SQLSelect("*", util.KeySystemUser, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&ret, q, nil)

	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *Service) GetByID(id uuid.UUID, addIfMissing bool) (*SystemUser, error) {
	ret := &SystemUser{}
	q := query.SQLSelect("*", util.KeySystemUser, "id = $1", "", 0, 0)
	err := s.db.Get(ret, q, nil, id)
	if err == sql.ErrNoRows {
		if addIfMissing {
			return s.new(id)
		}
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *Service) SaveProfile(prof *util.UserProfile) (*util.UserProfile, error) {
	s.logger.Info("updating user [" + prof.UserID.String() + "] from profile")
	cols := []string{"name", "role", "theme", "nav_color", "link_color", "picture", "locale"}
	q := query.SQLUpdate(util.KeySystemUser, cols, fmt.Sprintf("id = $%v", len(cols)+1))
	err := s.db.UpdateOne(q, nil, prof.Name, prof.Role.Key, prof.Theme.String(), prof.NavColor, prof.LinkColor, prof.Picture, prof.Locale.String(), prof.UserID)
	if err != nil {
		return nil, err
	}
	return prof, nil
}

func (s *Service) UpdateUserName(userID uuid.UUID, name string) error {
	s.logger.Info("updating user [" + userID.String() + "]")
	cols := []string{"name"}
	q := query.SQLUpdate(util.KeySystemUser, cols, fmt.Sprintf("id = $%v", len(cols)+1))
	err := s.db.UpdateOne(q, nil, name, userID)
	return err
}

func (s *Service) SetRole(userID uuid.UUID, role util.Role) error {
	s.logger.Info("updating user role [" + userID.String() + "]")
	cols := []string{"role"}
	q := query.SQLUpdate(util.KeySystemUser, cols, fmt.Sprintf("id = $%v", len(cols)+1))
	err := s.db.UpdateOne(q, nil, role.Key, userID)
	return err
}
