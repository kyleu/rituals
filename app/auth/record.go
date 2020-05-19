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
	_, err := s.db.Exec(q, r.ID, r.UserID, r.Provider.Key, r.ProviderID, r.Expires, r.Name, r.Email, r.Picture)
	if err != nil {
		return nil, err
	}

	return s.GetByID(r.ID)
}

func (s *Service) UpdateRecord(r *Record) error {
	q := "update auth set expires = $1, name = $2, email = $3, picture = $4 where id = $5"
	_, err := s.db.Exec(q, r.Expires, r.Name, r.Email, r.Picture, r.ID)
	return err
}

func (s *Service) List(params *query.Params) ([]*Record, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyAuth, params, &query.Ordering{Column: "created", Asc: false})
	var dtos []recordDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.KeyAuth, "", params.OrderByString(), params.Limit, params.Offset))
	if err != nil {
		return nil, err
	}
	return toRecords(dtos), nil
}

func (s *Service) GetByID(authID uuid.UUID) (*Record, error) {
	dto := &recordDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", util.KeyAuth, "id = $1", "", 0, 0), authID)
	if err != nil {
		return nil, err
	}
	return dto.ToRecord(), nil
}

func (s *Service) GetByProviderID(key string, code string) (*Record, error) {
	dto := &recordDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", util.KeyAuth, "provider = $1 and provider_id = $2", "", 0, 0), key, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToRecord(), nil
}

func (s *Service) GetByUserID(userID uuid.UUID, params *query.Params) ([]*Record, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyAuth, params, &query.Ordering{Column: "created", Asc: false})
	var dtos []recordDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.KeyAuth, "user_id = $1", params.OrderByString(), params.Limit, params.Offset), userID)
	if err != nil {
		return nil, err
	}
	return toRecords(dtos), nil
}

func toRecords(dtos []recordDTO) []*Record {
	ret := make([]*Record, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToRecord())
	}
	return ret
}
