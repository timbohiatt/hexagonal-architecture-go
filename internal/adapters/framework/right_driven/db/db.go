package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// Connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connction falure: %v", err)
	}
	// Test Connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v")
	}
	// Success
	return &Adapter{db: db}, nil
}

func (dba Adapter) CloseDbConnection() error {
	err := dba.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
	return nil
}

func (dba Adapter) AddToHistory(answer int32, operation string) error {

	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		log.Fatalf("db query failure: %v", err)
	}

	_, err = dba.db.Exec(queryString, args...)
	if err != nil {
		log.Fatalf("db exec failure: %v", err)
	}

	return nil
}
