package auth

import (
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
)

func (s *Service) NewRecord(r *Record) (*Record, error) {
	q := `insert into auth (id, user_id, k, v, expires, name, email, picture) values (
    $1, $2, $3, $4, $5, $6, $7, $8
	)`
	_, err := s.db.Exec(q, r.ID, r.UserID, r.K, r.V, r.Expires, r.Name, r.Email, r.Picture)
	if err != nil {
		return nil, err
	}

	return s.GetByID(r.ID)
}

func (s *Service) List() ([]*Record, error) {
	var dtos []recordDTO
	q := "select * from auth order by created desc"
	err := s.db.Select(&dtos, q)
	if err != nil {
		return nil, err
	}
	ret := make([]*Record, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToRecord())
	}
	return ret, nil
}

func (s *Service) GetByID(authID uuid.UUID) (*Record, error) {
	dto := &recordDTO{}
	err := s.db.Get(dto, "select * from auth where id = $1", authID)
	if err != nil {
		return nil, err
	}
	return dto.ToRecord(), nil
}

func (s *Service) GetByKV(key string, code string) (*Record, error) {
	dto := &recordDTO{}
	err := s.db.Get(dto, "select * from auth where k = $1 and v = $2", key, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToRecord(), nil
}

func (s *Service) GetByUserID(userID uuid.UUID, limit int) ([]*Record, error) {
	var dtos []recordDTO
	q := "select * from auth where user_id = $1 order by created desc"
	if limit > 0 {
		q += fmt.Sprint(" limit ", limit)
	}
	err := s.db.Select(&dtos, q, userID)
	if err != nil {
		return nil, err
	}
	ret := make([]*Record, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToRecord())
	}
	return ret, nil
}
