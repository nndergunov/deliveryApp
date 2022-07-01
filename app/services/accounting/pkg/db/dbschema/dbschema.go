// Package dbschema contains the database schema, migrations and seeding data.
package dbschema

import (
	"context"
	"database/sql"
	_ "embed" // Calls init function.
	"fmt"

	"github.com/ardanlabs/darwin"

	database "github.com/nndergunov/deliveryApp/app/services/accounting/pkg/db"
)

var (
	//go:embed sql/schema.sql
	schemaDoc string

	//go:embed sql/seed.sql
	seedDoc string

	//go:embed sql/delete.sql
	deleteDoc string
)

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(ctx context.Context, db *sql.DB) error {
	if err := database.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	driver, err := darwin.NewGenericDriver(db, darwin.PostgresDialect{})
	if err != nil {
		return fmt.Errorf("construct darwin driver: %w", err)
	}

	d := darwin.New(driver, darwin.ParseMigrations(schemaDoc))
	return d.Migrate()
}

// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func Seed(ctx context.Context, db *sql.DB) error {
	if err := database.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(seedDoc); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

// DeleteAll runs the set of Drop-table queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func DeleteAll(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(deleteDoc); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
