package iofs_test

import (
	"database/sql"
	"embed"
	"log"

	"github.com/vearutop/gooselite/iofs"
)

//go:embed testdata
var migrations embed.FS

func ExampleUp() {
	db, err := sql.Open("sqlite3", "sql.db")
	if err != nil {
		log.Fatalf("Failed to open test database: %v", err)
	}

	if err := iofs.Up(db, migrations, "testdata/migrations"); err != nil {
		log.Fatalf("Failed to run up migrations: %v", err)
	}
}
