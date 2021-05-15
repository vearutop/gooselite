package iofs

import (
	"database/sql"
	"io/fs"

	goose "github.com/vearutop/gooselite"
)

// Redo rolls back the most recently applied migration, then runs it again.
func Redo(db *sql.DB, fsys fs.FS, dir string) error {
	migrations, err := CollectMigrations(fsys, dir, 0, goose.MaxVersion)
	if err != nil {
		return err
	}

	return migrations.Redo(db)
}
