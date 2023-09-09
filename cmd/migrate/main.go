package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"os"
)

var (
	migrationPath = "./db/migrations"
	dsn           = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", "doo-dev", "123456", "localhost", "online-chat")
)

type migrator struct {
	migrations *migrate.FileMigrationSource
	dialect    string
}

func newMigrations() *migrator {
	migrations := &migrate.FileMigrationSource{Dir: migrationPath}
	return &migrator{migrations: migrations, dialect: "postgres"}
}

func (m migrator) up() {
	db, err := sql.Open(m.dialect, dsn)
	if err != nil {
		panic(fmt.Errorf("cant open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't apply migrations: %v", err))
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
func (m migrator) down() {
	db, err := sql.Open(m.dialect, dsn)
	if err != nil {
		panic(fmt.Errorf("cant open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't apply migrations: %v", err))

	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func main() {
	cmd := os.Args[1]
	m := newMigrations()

	switch cmd {
	case "up":
		m.up()
	case "down":
		m.down()
	default:
		fmt.Println("wrong command on migrations. provided args are `up` `down`")
	}
}
