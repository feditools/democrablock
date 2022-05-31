package migrations

import (
	migrate "github.com/tyrm/immudb-migrate"
)

// Migrations provides migrations for bun.
var Migrations = migrate.NewMigrations()
