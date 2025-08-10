package db

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectPostgre(dbURL string, agg interface{}) (*gorm.DB, error) {

	dsn, err := urlToDsn(dbURL)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = createSchema(dbURL, db)
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(agg); err != nil {
		return nil, err
	}

	return db, nil
}

func urlToDsn(dbURL string) (string, error) {
	u, err := url.Parse(dbURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse database URL: %w", err)
	}

	q := u.Query()
	sslmode := q.Get("sslmode")
	if sslmode == "" {
		sslmode = "disable"
	}
	timezone := q.Get("TimeZone")
	if timezone == "" {
		timezone = "UTC"
	}
	schema := q.Get("schema")
	if schema == "" {
		schema = "public"
	}

	password, _ := u.User.Password()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=%s TimeZone=%s",
		u.Hostname(),
		u.User.Username(),
		password,
		u.Path[1:],
		u.Port(),
		schema,
		sslmode,
		timezone,
	)

	return dsn, nil
}

func createSchema(dbURL string, db *gorm.DB) error {

	schema, err := urlToSchema(dbURL)
	if err != nil {
		return err
	}

	err = db.Exec(`CREATE SCHEMA IF NOT EXISTS "` + schema + `"`).Error
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	err = db.Exec(`SET search_path TO "` + schema + `"`).Error
	if err != nil {
		return fmt.Errorf("failed to set search_path: %w", err)
	}

	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		return fmt.Errorf("failed to create uuid-ossp extension: %w", err)
	}

	return nil
}

func urlToSchema(dbURL string) (string, error) {
	u, err := url.Parse(dbURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse database URL: %w", err)
	}

	q := u.Query()
	schema := q.Get("schema")
	if schema == "" {
		schema = "public"
	}

	return schema, nil
}
