package migrations

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
    migrationsPath string
	dbUrl 		   string
}

func NewMigrator (migrationPath string, dbURL string) *Migrator { 
	return &Migrator{
		migrationsPath: migrationPath,
		dbUrl: dbURL,
	}
}

func (migrator *Migrator) RunMigration() error {
	migrate, err := migrate.New(migrator.migrationsPath, migrator.dbUrl)
	if err != nil {
		fmt.Println(err.Error())
        return err
    }

	if err := migrate.Up(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}