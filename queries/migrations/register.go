package migrations

import "github.com/kyleu/rituals/app/lib/database/migrate"

func LoadMigrations(debug bool) {
	migrate.RegisterMigration("create initial database", Migration1InitialDatabase(debug))
}
