package postgreDatabase

import (
	"context"
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"

	"os"

	"github.com/uptrace/bun/migrate"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	// postgres driver
	_ "github.com/lib/pq"
)

var instance *bun.DB

var Migrations = migrate.NewMigrations()

var ctx = context.Background()

// GetDB returns the database single instance
func GetDB() *bun.DB {

	if instance == nil {
		log.Info("creating instance of repository")
		instance = initDB()

		doMigrations := os.Getenv("DO_MIGRATIONS")
		if doMigrations == "true" {
			migrateDB()
		}

	}
	return instance
}

func initDB() *bun.DB {

	log.Info("connecting to db")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))

	db := bun.NewDB(sqldb, pgdialect.New())
	log.Info("connected to db")
	return db
}
func migrateDB() {
	if err := Migrations.DiscoverCaller(); err != nil {
		log.Info("discover caller")
		panic(err)
	}
	migrator := migrate.NewMigrator(instance, Migrations)
	err := migrator.Init(ctx)
	if err != nil {
		log.Errorf("error with init migrator, err: %v", err)
		panic(err)
	}
	group, err := migrator.Migrate(ctx)
	if err != nil {
		log.Errorf("error with migrate migrator, err: %v", err)
		panic(err)
	}
	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
	}
	fmt.Printf("migrated to %s\n", group)
}

// Close closes the database connection
func Close() {
	instance.Close()
}
