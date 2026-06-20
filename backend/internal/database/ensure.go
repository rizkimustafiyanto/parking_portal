package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"backend/internal/config"
)

func EnsureDatabase(cfg *config.Config) error {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool

	query := `
	SELECT EXISTS (
		SELECT 1
		FROM pg_database
		WHERE datname = $1
	)
	`

	err = db.QueryRow(query, cfg.DBName).Scan(&exists)

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	_, err = db.Exec(
		fmt.Sprintf(
			`CREATE DATABASE "%s"`,
			cfg.DBName,
		),
	)

	return err
}