package cmd

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/database/migrate"
	"github.com/kyleu/rituals/app/lib/log"
	"github.com/kyleu/rituals/queries/migrations"
)

func migrateCmd() *cobra.Command {
	f := func(*cobra.Command, []string) error { return runMigrations(rootCtx) }
	ret := newCmd("migrate", "Runs database migrations and exits", f)
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
