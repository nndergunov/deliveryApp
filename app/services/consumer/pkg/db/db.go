package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenDB(driver, url string) (*sql.DB, error) {
	database, err := sql.Open(driver, url)
	if err != nil {
		return nil, fmt.Errorf("OpenDB: %w", err)
	}
	return database, nil
}
