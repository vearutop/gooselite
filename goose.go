package gooselite

import (
	"fmt"
)

var (
	minVersion      = int64(0)
	maxVersion      = int64((1 << 63) - 1)
	timestampFormat = "20060102150405"
	verbose         = false
)

// SetVerbose set the goose verbosity mode.
func SetVerbose(v bool) {
	verbose = v
}

// Run runs a goose command.
func Run(command string, dir string, args ...string) error {
	switch command {
	case "create":
		if len(args) == 0 {
			return fmt.Errorf("create must be of form: goose [OPTIONS] DRIVER DBSTRING create NAME [go|sql]")
		}

		migrationType := "go"

		if len(args) == 2 {
			migrationType = args[1]
		}

		if err := Create(dir, args[0], migrationType); err != nil {
			return err
		}
	case "fix":
		if err := Fix(dir); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%q: no such command", command)
	}

	return nil
}
