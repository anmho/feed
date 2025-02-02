package test

import (
	"context"
	"database/sql"
	"log"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for database/sql
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)



func MakeLocalDB(t *testing.T) *sql.DB {
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	sqlDir := filepath.Join(currentDir, "../../sql/schema.sql")
	
	ctx := context.Background()
	pgContainer, err := postgres.Run(
		ctx,
		"postgres:15.6",
		// postgres.WithInitScripts(sqlDir),
		postgres.WithInitScripts(sqlDir),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
		// should add init scripts here
	)
	
	require.NoError(t, err, "creating test container")
	require.NotNil(t, pgContainer)
	require.NoError(t, pgContainer.Start(ctx), "starting test container")
	log.Println("running", pgContainer.IsRunning())
	dbURL := pgContainer.MustConnectionString(ctx)
	db, err := sql.Open("pgx", dbURL)
	require.NoError(t, err, "opening test db")
	require.NoError(t, db.PingContext(ctx), "pinging test db")

	t.Cleanup(func() {
		duration := time.Second * 30
		require.NoError(t, pgContainer.Stop(ctx, &duration), "stopping test db")
	})


	// Return the db instance and a cleanup function
	return db
}
