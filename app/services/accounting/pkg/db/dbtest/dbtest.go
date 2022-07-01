package dbtest

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"os"
	"testing"

	database "github.com/nndergunov/deliveryApp/app/services/accounting/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/db/dbschema"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/docker"
)

// StartDB starts a database instance.
func StartDB() (*docker.Container, error) {
	image := "postgres:14-alpine"
	port := "5432"
	args := []string{"-e", "POSTGRES_PASSWORD=postgres"}

	return docker.StartContainer(image, port, args...)
}

// StopDB stops a running database instance.
func StopDB(c *docker.Container) {
	docker.StopContainer(c.ID)
}

// NewUnit creates a test database inside a Docker container. It creates the
// required table structure but the database is otherwise empty. It returns
// the database to use as well as a function to call at the end of the test.
func NewUnit(t *testing.T, c *docker.Container, dbName string) (*logger.Logger, *sql.DB, func()) {
	ctx := context.TODO()
	//defer cancel()

	var dbURL = fmt.Sprintf(
		"host=" + "host.docker.internal" +
			" user=postgres" +
			" password=postgres" +
			" dbname=postgres" +
			" sslmode=true")

	dbM, err := database.OpenDB("postgres", dbURL)
	if err != nil {
		t.Fatalf("Opening database connection: %v", err)
	}

	t.Log("Waiting for database to be ready ...")

	if err := database.StatusCheck(ctx, dbM); err != nil {
		t.Fatalf("status check database: %v", err)
	}

	t.Log("Database ready")

	if _, err := dbM.ExecContext(context.Background(), "CREATE DATABASE "+dbName); err != nil {
		t.Fatalf("creating database %s: %v", dbName, err)
	}
	dbM.Close()

	// =========================================================================
	db, err := database.OpenDB("postgres", fmt.Sprintf(
		"host="+"host.docker.internal"+
			" user=postgres"+
			" password=postgres"+
			" dbname="+dbName+
			" sslmode=true"))
	if err != nil {
		t.Fatalf("Opening database connection: %v", err)
	}

	t.Log("seed database ...")

	if err := dbschema.Seed(ctx, db); err != nil {
		docker.DumpContainerLogs(t, c.ID)
		t.Fatalf("Seeding error: %s", err)
	}

	t.Log("Ready for testing ...")

	log := logger.NewLogger(os.Stdout, "service: ")
	// teardown is the function that should be invoked when the caller is done
	// with the database.
	teardown := func() {
		t.Helper()
		db.Close()
	}

	return log, db, teardown
}
