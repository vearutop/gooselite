package gooselite

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const seqVersionTemplate = "%05v"

// Fix renames timestamped filenames with sequential versions.
func Fix(dir string) error {
	migrations, err := CollectMigrations(dir, minVersion, maxVersion)
	if err != nil {
		return err
	}

	// split into timestamped and versioned migrations
	tsMigrations := migrations.timestamped()

	vMigrations := migrations.versioned()

	// Initial version.
	version := int64(1)
	if last, err := vMigrations.Last(); err == nil {
		version = last.Version + 1
	}

	// fix filenames by replacing timestamps with sequential versions
	for _, tsm := range tsMigrations {
		oldPath := tsm.Source
		newPath := strings.Replace(
			oldPath,
			fmt.Sprintf("%d", tsm.Version),
			fmt.Sprintf(seqVersionTemplate, version),
			1,
		)

		if err := os.Rename(oldPath, newPath); err != nil {
			return err
		}

		log.Printf("RENAMED %s => %s", filepath.Base(oldPath), filepath.Base(newPath))
		version++
	}

	return nil
}
