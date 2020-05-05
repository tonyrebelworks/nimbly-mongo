package mysql

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db       *sqlx.DB
	err      error
	dbDriver = "mysql"
)

// Connect ...
func Connect(dsn string) *sqlx.DB {
	db := sqlx.MustConnect(dbDriver, dsn)
	return db
}

// Close close the DB connection
func Close() error {
	return db.Close()
}
