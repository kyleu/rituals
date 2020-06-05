package database

import (
	"fmt"
	"logur.dev/logur"
	"time"
)

func DBWipe(s *Service, logger logur.Logger) error {
	for _, file := range initialSchemaMigrations {
		_, err := exec(file, s, logger)
		if err != nil {
			return err
		}
	}
	return nil
}

func Migrate(s *Service) error {
	maxIdx := maxMigrationIdx(s)
	// s.logger.Info(fmt.Sprintf("migrating database schema: %v", maxIdx))

	for i, file := range databaseMigrations {
		idx := i + 1
		if (idx) > maxIdx {
			s.logger.Info(fmt.Sprintf("applying database migration [%v]: %v", idx, file.Name))
			sql, err := exec(file, s, s.logger)
			if err != nil {
				return err
			}
			err = newMigration(s, Migration{
				Idx:     idx,
				Title:   file.Name,
				Src:     sql,
				Created: time.Time{},
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
