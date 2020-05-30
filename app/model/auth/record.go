package auth

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
)

func (s *Service) NewRecord(r *Record) (*Record, error) {
	q := query.SQLInsert(util.KeyAuth, []string{util.KeyID, util.WithDBID(util.KeyUser), util.KeyProvider, util.WithDBID(util.KeyProvider), "user_list_id", "user_list_name", "access_token", "expires", util.KeyName, util.KeyEmail, "picture"}, 1)
	err := s.db.Insert(q, nil, r.ID, r.UserID, r.Provider.Key, r.ProviderID, r.UserListID, r.UserListName, r.AccessToken, r.Expires, r.Name, r.Email, r.Picture)
	if err != nil {
		return nil, err
	}

	return s.GetByID(r.ID), nil
}

func (s *Service) UpdateRecord(r *Record) error {
	cols := []string{"user_list_id", "user_list_name", "access_token", "expires", util.KeyName, util.KeyEmail, "picture"}
	q := query.SQLUpdate(util.KeyAuth, cols, fmt.Sprintf("%v = $%v", util.KeyID, len(cols)+1))
	return s.db.UpdateOne(q, nil, r.UserListID, r.UserListName, r.AccessToken, r.Expires, r.Name, r.Email, r.Picture, r.ID)
}

func (s *Service) List(params *query.Params) Records {
	params = query.ParamsWithDefaultOrdering(util.KeyAuth, params, query.DefaultCreatedOrdering...)
	var dtos []recordDTO
	q := query.SQLSelect("*", util.KeyAuth, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		util.LogError(s.logger, "error retrieving auth records: %+v", err)
		return nil
	}
	return toRecords(dtos)
}

func (s *Service) GetByID(authID uuid.UUID) *Record {
	dto := &recordDTO{}
	q := query.SQLSelectSimple("*", util.KeyAuth, util.KeyID + " = $1")
	err := s.db.Get(dto, q, nil, authID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		util.LogError(s.logger, "error getting auth record by id [%v]: %+v", authID, err)
		return nil
	}
	return dto.ToRecord()
}

func (s *Service) GetByProviderID(key string, code string) *Record {
	dto := &recordDTO{}
	q := query.SQLSelectSimple("*", util.KeyAuth, "provider = $1 and provider_id = $2")
	err := s.db.Get(dto, q, nil, key, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		util.LogError(s.logger, "error getting auth record by provider [%v:%v]: %+v", key, code, err)
		return nil
	}
	return dto.ToRecord()
}

func (s *Service) GetByUserID(userID uuid.UUID, params *query.Params) Records {
	params = query.ParamsWithDefaultOrdering(util.KeyAuth, params, query.DefaultCreatedOrdering...)
	var dtos []recordDTO
	q := query.SQLSelect("*", util.KeyAuth, "user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		util.LogError(s.logger, "error retrieving auths for user [%v]: %+v", userID, err)
		return nil
	}
	return toRecords(dtos)
}

func (s *Service) Delete(authID uuid.UUID) error {
	q := query.SQLDelete(util.KeyAuth, util.KeyID + " = $1")
	return s.db.DeleteOne(q, nil, authID)
}

func toRecords(dtos []recordDTO) Records {
	ret := make(Records, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToRecord())
	}
	return ret
}
