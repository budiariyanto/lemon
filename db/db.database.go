package db

import (
	"database/sql"
)

type Database interface {
	Connect() (*sql.DB, error)
}

func NewDbConnection(db Database) (*sql.DB, error) {
	return db.Connect()
}
