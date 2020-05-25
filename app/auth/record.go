package auth

import (
	"database/sql"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/query"
)

func (s *Service) NewRecord(r *Record) (*Record, error) {
	q := `insert into auth (id, user_id, provider, provider_id, expires, name, email, picture) values (
    $1, $2, $3, $4, $5, $6, $7, $8
	)`
	err := s.db.Insert(q, nil, r.ID, r.UserID, r.Provider.Key, r.ProviderID, r.Expires, r.Name, r.Email, r.Picture)
	if err != nil {
		return nil, err
	}

	return s.GetByID(r.ID)
}

func (s *Service) UpdateRecord(r *Record) error {
	q := "update auth set expires = $1, name = $2, email = $3, picture = $4 where id = $5"
	return s.db.UpdateOne(q, nil, r.Expires, r.Name, r.Email, r.Picture, r.ID)
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
	q := "delete from auth where id = $1"
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
