package db

import (
	"time"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"battleship/app/errors"
)

const (
	PgTimestampFormatString = "2006-01-02 15:04:05.000000" // Format string for Time.Format() method
)

type pqDatabase struct {
	db *sqlx.DB
}

func NewPQDatabase(dsn string, maxIdleConns, maxOpenConns int, connMaxLifetime time.Duration) (Storage, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	if err := db.Ping(); err != nil {
		return nil, errors.NewDatabaseError(errors.DatabaseConnectionError, err)
	}

	return &pqDatabase{db: db}, nil
}

func (p *pqDatabase) Close() error {
	return p.db.Close()
}

func (p *pqDatabase) BeginTx() (StorageTx, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, err
	}

	return &PQDatabaseITx{tx: tx}, nil
}

func (p *pqDatabase) Exec(query string, params ...interface{}) error {
	_, err := p.db.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (p *pqDatabase) Query(result interface{}, query string, params ...interface{}) error {
	return p.db.Select(result, query, params...)
}

type PQDatabaseITx struct {
	tx *sqlx.Tx
}

func (p *PQDatabaseITx) BeginTx() (StorageTx, error) {
	return p, nil
}

func (p *PQDatabaseITx) Exec(query string, params ...interface{}) error {
	_, err := p.tx.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (p *PQDatabaseITx) Query(result interface{}, query string, params ...interface{}) error {
	return p.tx.Select(result, query, params...)
}

func (p *PQDatabaseITx) Commit() error {
	return p.tx.Commit()
}

func (p *PQDatabaseITx) Rollback() error {
	return p.tx.Rollback()
}
