package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func OpenDatabase(cfg PostgresConfig) error {
	var err error

	connStr := fmt.Sprintf("user=%s port=%s password=%s dbname=%s sslmode=disable",
		cfg.Username,
		cfg.Port,
		cfg.Password,
		cfg.DBName,
	)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

func CloseDatabase() error {
	return db.Close()
}
