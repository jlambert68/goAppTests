package server

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DbPool *pgxpool.Pool

func ConnectToDB() {
	dbConnectionURL := "postgres://testuser:password@127.0.0.1:5432/testdb"

	var err error

	DbPool, err = pgxpool.Connect(context.Background(), dbConnectionURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer dbpool.Close()

	var greeting string
	err = DbPool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
