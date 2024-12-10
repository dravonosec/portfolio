package database

import (
	"context"
	"fmt"
	"musicLibrary/internal/config"
	"os"

	"github.com/jackc/pgx/v5"
)

type DatabaseCreator struct {
	databaseUrl string
}

func NewDatabase(config *config.Config) *DatabaseCreator {
    return &DatabaseCreator{databaseUrl: config.GetDBConnString()}
}

func (db *DatabaseCreator) Create(databaseName string) error { 
	conn, err := pgx.Connect(context.Background(), db.databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	
	_, err = conn.Query(context.Background(), "CREATE DATABASE " + databaseName + ";")
	if (err != nil) {
		fmt.Println(err.Error())
		return err
	}

	defer conn.Close(context.Background())
	return nil
}