package database

import (
	"Labs2/internal/config"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Connection(config config.DB) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	queries := []string{
		`CREATE TABLE IF NOT EXISTS jokes (
            id SERIAL PRIMARY KEY,
            joke_id VARCHAR(255) UNIQUE,
            category VARCHAR(100),
            joke TEXT,
            created_at TIMESTAMP DEFAULT NOW()
        );`,
		`CREATE TABLE IF NOT EXISTS cats (
            id SERIAL PRIMARY KEY,
            cat_id VARCHAR(255) UNIQUE,
            url TEXT,
            created_at TIMESTAMP DEFAULT NOW()
        );
		CREATE TABLE IF NOT EXISTS number_facts (
        	id SERIAL PRIMARY KEY,
        	num INT,
        	fact TEXT,
        	created_at TIMESTAMP DEFAULT NOW()
    	);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func MigrateUp(config config.DB) error {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Name)
	m, err := migrate.New(
		"file://"+config.MigrationsPath,
		connectionString)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
