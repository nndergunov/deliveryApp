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

//go:embed sql/schema.sql
var schemaDoc string

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
