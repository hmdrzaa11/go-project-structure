package database

import (
	"database/sql"
	"time"
)

const maxOpenConnection = 10
const maxIdleConnection = 5
const maxConnectionLifetime = 5 * time.Minute

type Database struct {
	Client *sql.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	//test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConnection)
	db.SetConnMaxLifetime(maxConnectionLifetime)
	//after configure we also test

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}
