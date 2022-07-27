package dbtest

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	database "github.com/nndergunov/deliveryApp/app/services/consumer/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/db/dbschema"
	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/docker"
)

// StartDB starts a database instance.
func StartDB() (*docker.Container, error) {
	image := "postgres:latest"
	port := "5000"
	args := []string{"-e", "POSTGRES_PASSWORD=postgres"}

	return docker.StartContainer(image, port, args...)
}

// StopDB stops a running database instance.
func StopDB(c *docker.Container) {
	if err := docker.StopContainer(c.ID); err != nil {
		log.Println(err)
	}
}

// NewUnit creates a test database inside a Docker container. It creates the
// required table structure but the database is otherwise empty. It returns
// the database to use as well as a function to call at the end of the test.
func NewUnit(t *testing.T, c *docker.Container, dbName string) (*sql.DB, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	dbM, err := database.OpenDB("postgres",
		fmt.Sprintf("host="+c.Host+
			" port="+c.Port+
			" user=postgres"+
			" password=postgres"+
			" dbname=postgres"+
			" sslmode=disable"))
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

	db, err := database.OpenDB("postgres",
		fmt.Sprintf("host="+c.Host+
			" port="+c.Port+
			" user=postgres"+
			" password=postgres"+
			" dbname="+dbName+
			" sslmode=disable"))
	if err != nil {
		t.Fatalf("Opening database connection: %v", err)
	}

	t.Log("Migrate database ...")

	if err := dbschema.Migrate(ctx, db); err != nil {
		docker.DumpContainerLogs(t, c.ID)
		t.Fatalf("Migrating error: %s", err)
	}

	t.Log("Ready for testing ...")

	// teardown is the function that should be invoked when the caller is done
	// with the database.
	teardown := func() {
		t.Helper()
		db.Close()
	}

	return db, teardown
}
