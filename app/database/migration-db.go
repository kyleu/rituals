package database

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) ListMigrations(params *query.Params) Migrations {
	params = query.ParamsWithDefaultOrdering(util.KeyMigration, params, query.DefaultCreatedOrdering...)

	var dtos []migrationDTO
	q := query.SQLSelect("*", util.KeyMigration, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.Select(&dtos, q, nil)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving migrations: %+v", err))
		return nil
	}

	return toMigrations(dtos)
}

func (s *Service) GetMigrationByIdx(idx int) *Migration {
	var dto = &migrationDTO{}
	q := query.SQLSelectSimple("*", util.KeyMigration, "idx = $1")
	err := s.Get(dto, q, nil, idx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Error(fmt.Sprintf("error getting migration by idx [%v]: %+v", idx, err))
		return nil
	}
	return dto.toMigration()
}

func newMigration(s *Service, e Migration) error {
	q := query.SQLInsert(util.KeyMigration, []string{util.KeyIdx, util.KeyTitle, "src"}, 1)
	return s.Insert(q, nil, e.Idx, e.Title, e.Src)
}

func maxMigrationIdx(s *Service) int {
	q := query.SQLSelectSimple("max(idx) as x", util.KeyMigration, "")
	max, err := s.SingleInt(q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting migrations: %+v", err))
		return -1
	}
	return int(max)
}

func toMigrations(dtos []migrationDTO) Migrations {
	ret := make(Migrations, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.toMigration())
	}

	return ret
}
