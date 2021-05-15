# gooselite

[![Build Status](https://github.com/vearutop/gooselite/workflows/test-unit/badge.svg)](https://github.com/vearutop/gooselite/actions?query=branch%3Amaster+workflow%3Atest-unit)
[![Coverage Status](https://codecov.io/gh/vearutop/gooselite/branch/master/graph/badge.svg)](https://codecov.io/gh/vearutop/gooselite)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/vearutop/gooselite)
[![Time Tracker](https://wakatime.com/badge/github/vearutop/gooselite.svg)](https://wakatime.com/badge/github/vearutop/gooselite)
![Code lines](https://sloc.xyz/github/vearutop/gooselite/?category=code)
![Comments](https://sloc.xyz/github/vearutop/gooselite/?category=comments)

`gooselite` is a reduced version of [`goose`](https://github.com/pressly/goose) tailored for embedding in your
application.

## Usage

Install `gooselite` binary with
```
go install github.com/vearutop/gooselite/cmd/gooselite
```

or download it from [releases](https://github.com/vearutop/gooselite/releases/).

Create migrations with `gooselite create add_some_table sql` in your migrations directory.

Apply migrations from your application.

```go
package main

import (
	"database/sql"
	"embed"
	"log"

	"github.com/vearutop/gooselite/iofs"
)

//go:embed migrations
var migrations embed.FS

func main() {
	db, err := sql.Open("sqlite3", "sql.db") // Open connection to your DB.
	if err != nil {
		log.Fatalf("Failed to open test database: %v", err)
	}

	// Apply migrations.
	if err := iofs.Up(db, migrations, "migrations"); err != nil {
		log.Fatalf("Failed to run up migrations: %v", err)
	}
}

```