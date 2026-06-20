package database

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

const migrationVersionTable = "schema_migrations"

type migrationFile struct {
	Version int
	Path    string
	Name    string
}

func Migrate(db *gorm.DB, migrationsDir string) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err := ensureMigrationTable(sqlDB); err != nil {
		return err
	}

	applied, err := loadAppliedVersions(sqlDB)
	if err != nil {
		return err
	}

	files, err := loadMigrationFiles(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if applied[file.Version] {
			continue
		}

		if err := applyMigration(sqlDB, file); err != nil {
			return err
		}
	}

	return nil
}

func ensureMigrationTable(db *sql.DB) error {
	_, err := db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			version INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`, migrationVersionTable))
	return err
}

func loadAppliedVersions(db *sql.DB) (map[int]bool, error) {
	rows, err := db.Query(fmt.Sprintf(`SELECT version FROM %s`, migrationVersionTable))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applied := make(map[int]bool)
	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		applied[version] = true
	}

	return applied, rows.Err()
}

func loadMigrationFiles(dir string) ([]migrationFile, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	files := make([]migrationFile, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if filepath.Ext(name) != ".sql" {
			continue
		}

		base := strings.TrimSuffix(name, ".sql")
		parts := strings.SplitN(base, "_", 2)
		if len(parts) == 0 {
			continue
		}

		version, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		files = append(files, migrationFile{
			Version: version,
			Path:    filepath.Join(dir, name),
			Name:    name,
		})
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].Version == files[j].Version {
			return files[i].Name < files[j].Name
		}
		return files[i].Version < files[j].Version
	})

	return files, nil
}

func applyMigration(db *sql.DB, file migrationFile) error {
	content, err := os.ReadFile(file.Path)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(string(content)); err != nil {
		_ = tx.Rollback()
		return err
	}

	if _, err := tx.Exec(
		fmt.Sprintf(`INSERT INTO %s (version, name) VALUES ($1, $2)`, migrationVersionTable),
		file.Version,
		file.Name,
	); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
