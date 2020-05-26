package auth

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/query"
)

func (s *Service) NewRecord(r *Record) (*Record, error) {
	q := query.SQLInsert(util.KeyAuth, []string{"id", "user_id", "provider", "provider_id", "user_list_id", "user_list_name", "access_token", "expires", "name", "email", "picture"}, 1)
	err := s.db.Insert(q, nil, r.ID, r.UserID, r.Provider.Key, r.ProviderID, r.UserListID, r.UserListName, r.AccessToken, r.Expires, r.Name, r.Email, r.Picture)
	if err != nil {
		return nil, err
	}

	return s.GetByID(r.ID)
}

func (s *Service) UpdateRecord(r *Record) error {
	cols := []string{"user_list_id", "user_list_name", "access_token", "expires", "name", "email", "picture"}
	q := query.SQLUpdate(util.KeyAuth, cols, fmt.Sprintf("id = $%v", len(cols)+1))
	return s.db.UpdateOne(q, nil, r.UserListID, r.UserListName, r.AccessToken, r.Expires, r.Name, r.Email, r.Picture, r.ID)
}

func (s *Service) List(params *query.Params) (Records, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyAuth, params, query.DefaultCreatedOrdering...)
	var dtos []recordDTO
	q := query.SQLSelect("*", util.KeyAuth, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		return nil, err
	}
	return toRecords(dtos), nil
}

func (s *Service) GetByID(authID uuid.UUID) (*Record, error) {
	dto := &recordDTO{}
	q := query.SQLSelect("*", util.KeyAuth, "id = $1", "", 0, 0)
	err := s.db.Get(dto, q, nil, authID)
	if err != nil {
		return nil, err
	}
	return dto.ToRecord(), nil
}

func (s *Service) GetByProviderID(key string, code string) (*Record, error) {
	dto := &recordDTO{}
	q := query.SQLSelect("*", util.KeyAuth, "provider = $1 and provider_id = $2", "", 0, 0)
	err := s.db.Get(dto, q, nil, key, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToRecord(), nil
}

func (s *Service) GetByUserID(userID uuid.UUID, params *query.Params) (Records, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyAuth, params, query.DefaultCreatedOrdering...)
	var dtos []recordDTO
	q := query.SQLSelect("*", util.KeyAuth, "user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		return nil, err
	}
	return toRecords(dtos), nil
}

func (s *Service) Delete(authID uuid.UUID) error {
	q := query.SQLDelete(util.KeyAuth, "id = $1")
	_, err := s.db.Delete(q, nil, authID)
	return err
}

func toRecords(dtos []recordDTO) Records {
	ret := make(Records, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToRecord())
	}
	return ret
}
