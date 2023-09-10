// file: main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	// We are required to import the driver, although it's not used explicitly in our code.
	// To indicate this, we use the `_` symbol.
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// The function `sql.Open` creates a new `*sql.DB` instance. We specify the driver name
	// and the URI for our database. In this case, we're using a Postgres URI.
	db, err := sql.Open("pgx", "postgresql://localhost:5432/bird_encyclopedia")
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	// To confirm the connection to our database, we can use the `Ping` method.
	// If no error is returned, we can assume a successful connection.
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach the database: %v", err)
	}
	fmt.Println("database is accessible")
}
