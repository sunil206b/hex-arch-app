package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) *Adapter {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}
	return &Adapter{
		db: db,
	}
}

func (ad *Adapter) CloseDbConnection() {
	err := ad.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (ad *Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}
	_, err = ad.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
