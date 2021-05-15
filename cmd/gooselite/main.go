package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bool64/dev/version"
	goose "github.com/vearutop/gooselite"
)

var (
	flags      = flag.NewFlagSet("gooselite", flag.ExitOnError)
	dir        = flags.String("dir", ".", "directory with migration files")
	table      = flags.String("table", "goose_db_version", "migrations table name")
	verbose    = flags.Bool("v", false, "enable verbose mode")
	help       = flags.Bool("h", false, "print help")
	ver        = flags.Bool("version", false, "print version")
	sequential = flags.Bool("s", false, "use sequential numbering for new migrations")
)

func main() {
	flags.Usage = usage

	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if *ver {
		fmt.Println(version.Info().Version)

		return
	}

	if *verbose {
		goose.SetVerbose(true)
	}

	if *sequential {
		goose.SetSequential(true)
	}

	goose.SetTableName(*table)

	args := flags.Args()
	if len(args) == 0 || *help {
		flags.Usage()

		return
	}

	switch args[0] {
	case "create":
		if err := goose.Run("create", *dir, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}

		return
	case "fix":
		if err := goose.Run("fix", *dir); err != nil {
			log.Fatalf("goose run: %v", err)
		}

		return
	}

	args = mergeArgs(args)
	if len(args) < 3 {
		flags.Usage()

		return
	}

	_, _, command := args[0], args[1], args[2]

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, *dir, arguments...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}

const (
	envGooseDriver   = "GOOSE_DRIVER"
	envGooseDBString = "GOOSE_DBSTRING"
)

func mergeArgs(args []string) []string {
	if len(args) < 1 {
		return args
	}

	if d := os.Getenv(envGooseDriver); d != "" {
		args = append([]string{d}, args...)
	}

	if d := os.Getenv(envGooseDBString); d != "" {
		args = append([]string{args[0], d}, args[1:]...)
	}

	return args
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix = `Usage: gooselite [OPTIONS] DRIVER DBSTRING COMMAND

or

Set environment key
GOOSE_DRIVER=DRIVER

Usage: goose [OPTIONS] COMMAND

Drivers:
    postgres
    mysql
    sqlite3
    mssql
    redshift
    clickhouse

Examples:
    goose sqlite3 ./foo.db create init sql
    goose sqlite3 ./foo.db create add_some_column sql
    goose sqlite3 ./foo.db create fetch_user_data go

Options:
`

	usageCommands = `
Commands:
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)
