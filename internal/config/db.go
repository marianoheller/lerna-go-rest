package config

import (
	"database/sql"
	"fmt"
	embedded "learn"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetDBConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	return db
}

func InitialiseDatabase() {
	db := GetDBConnection()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}

	d, err := iofs.New(embedded.MigrationsFS, "db/migrations")
	if err != nil {
		fmt.Println(err)
		return
	}

	m, err := migrate.NewWithInstance("iofs", d, "sqlite3", driver)

	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file://"+util.PathRelProjectDir("/db/migrations"),
	// 	"sqlite3",
	// 	driver,
	// )
	if err != nil {
		panic(err)
	}

	m.Log = &MigrationLogger{}

	if err := m.Up(); err != nil {
		panic(err)
	}
}

type MigrationLogger struct{}

func (*MigrationLogger) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
func (*MigrationLogger) Verbose() bool {
	return true
}
