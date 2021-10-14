package server

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

var DbPool *pgxpool.Pool

func ConnectToDB() {

	// Postgres version 10
	//dbConnectionURL := "postgres://testuser:password@127.0.0.1:5433/testdb"

	// Postgres version 14
	dbConnectionURL := "postgres://caxdbuser:password@127.0.0.1:5432/caxdb"

	var err error

	DbPool, err = pgxpool.Connect(context.Background(), dbConnectionURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer dbpool.Close()

	var version string
	err = DbPool.QueryRow(context.Background(), "SELECT VERSION()").Scan(&version)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(version)
}
