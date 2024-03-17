package embedded

import (
	"embed"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed db/migrations/*.sql
var MigrationsFS embed.FS
