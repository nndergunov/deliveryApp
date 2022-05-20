// Package db provides support for access the database.
package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Open knows how to open a database connection based on the configuration.
func Open(driver, dbUrl string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
