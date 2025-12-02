package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := sql.Open("postgres", "host=127.0.0.1 user=postgres password=postgres port=5432 sslmode=disable dbname=users")
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err.Error())
	}

	err = conn.Ping()
	if err != nil {
		log.Println("Ping failed", err.Error())
		return
	}

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Println("Instance failed", err.Error())
		return
	}
	m, err := migrate.NewWithDatabaseInstance("file://../cmd/migrate/migrations", "postgres", driver)
	if err != nil {
		log.Println("Migrate Instance failed", err.Error())
		return
	}

	// operation := flag.Arg(1)
	operation := flag.String("oper", "", "operation to perform")
	flag.Parse()
	fmt.Println("Opeartion is: ", *operation)
	if *operation == "up" {
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Println("No migration change detected!")
			} else {
				log.Fatalf("Migration up failed: %v", err)
			}

		}
		fmt.Println("Migration up completed successfully")
	}

	if *operation == "down" {
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Migration down failed: %v", err)
		}
		fmt.Println("Migration down completed successfully")
	}

}
