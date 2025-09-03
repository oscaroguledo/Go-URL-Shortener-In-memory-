package core

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
	}

	cmd := args[0]
	var arg string
	if len(args) > 1 {
		arg = args[1]
	}

	// TODO: Replace with actual settings initialization as needed
	// settings := config.Settings(
	// 	func(s *SettingStruct) {
	// 		s.DatabaseType = "sqlite3"
	// 		s.DatabaseURL = "urls.db"
	// 	},
	// )

	migrationsPath := "../migrations"
	dbURL := buildDBURL("sqlite3", "urls.db")

	switch cmd {
	case "new":
		if arg == "" {
			log.Fatal("Please provide a migration name (e.g., create_users_table).")
		}
		runCommand("migrate", "create", "-ext", "sql", "-dir", migrationsPath, arg)

	case "up":
		runCommand("migrate", "-database", dbURL, "-path", migrationsPath, "up", arg)

	case "down":
		runCommand("migrate", "-database", dbURL, "-path", migrationsPath, "down", arg)

	case "goto":
		if arg == "" {
			log.Fatal("Please provide a version number.")
		}
		runCommand("migrate", "-database", dbURL, "-path", migrationsPath, "goto", arg)

	case "version":
		runCommand("migrate", "-database", dbURL, "-path", migrationsPath, "version")

	case "drop":
		runCommand("migrate", "-database", dbURL, "-path", migrationsPath, "drop", "-f")

	default:
		usage()
	}
}

func runCommand(name string, args ...string) {
	// Clean args: remove empty strings
	var cleanArgs []string
	for _, a := range args {
		if a != "" {
			cleanArgs = append(cleanArgs, a)
		}
	}
	cmd := exec.Command(name, cleanArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Command failed: %v", err)
	}
}

func usage() {
	fmt.Println(`Usage: migrate <command> [n|version]

Commands:
  new <name>     – Create new migration
  up [n]         – Apply all or next n migrations
  down [n]       – Rollback last or last n migrations
  goto <ver>     – Migrate to a specific version
  version        – Print current migration version
  drop           – Drop everything (CAREFUL!)`)
	os.Exit(1)
}

func buildDBURL(dbType, dbPath string) string {
	switch dbType {
	case "sqlite3":
		// Ensure it's absolute path for sqlite
		abs, err := filepath.Abs(dbPath)
		if err != nil {
			log.Fatalf("Failed to get absolute path: %v", err)
		}
		return fmt.Sprintf("sqlite3://%s", abs)
	case "postgres", "mysql", "sqlserver", "mongodb":
		return dbPath
	default:
		log.Fatalf("Unsupported database type: %s", dbType)
	}
	return ""
}
