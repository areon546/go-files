package sqlite

import (
	"database/sql"

	"github.com/mattn/go-sqlite3"
)

// CreateNewSQLITEFile
// AddFile

type SQLiteFile struct {
	db *sql.DB
}

func Open(dsn string) (*SQLiteFile, error) {
	slite := &SQLiteFile{}

	sql.Register(dsn, &sqlite3.SQLiteDriver{})
	db, err := sql.Open(dsn, dsn)
	if err != nil {
		return nil, err
	}

	slite.db = db
	return slite, nil
}

func (slite *SQLiteFile) Query(query string, args ...any) (*sql.Rows, error) {
	return slite.db.Query(query, args...)
}

func (slite *SQLiteFile) Exec(query string, args ...any) (sql.Result, error) {
	return slite.db.Exec(query, args...)
}

func (slite *SQLiteFile) Prepare(query string) (*sql.Stmt, error) {
	return slite.db.Prepare(query)
}
