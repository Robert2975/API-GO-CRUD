package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// Connect initializes the database connection
func Connect() error {
	var err error
	DB, err = sqlx.Connect("postgres", "user=postgres password=123456 dbname=gin_crud sslmode=disable")
	if err != nil {
		return err
	}
	return nil
}
