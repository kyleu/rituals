package cmd

import (
	"context"

	"github.com/muesli/coral"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/database/migrate"
	"github.com/kyleu/rituals/app/lib/log"
	"github.com/kyleu/rituals/queries/migrations"
)

func migrateCmd() *coral.Command {
	f := func(*coral.Command, []string) error { return runMigrations(context.Background()) }
	ret := &coral.Command{Use: "migrate", Short: "Runs database migrations and exits", RunE: f}
	return ret
}

func runMigrations(ctx context.Context) error {
	logger, _ := log.InitLogging(false)
	db, err := database.OpenDefaultPostgres(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "unable to open database")
	}
	migrations.LoadMigrations(_flags.Debug)
	err = migrate.Migrate(ctx, db, logger)
	if err != nil {
		return errors.Wrap(err, "unable to run database migrations")
	}
	return nil
}
