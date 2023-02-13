package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresConn struct {
	db *sql.DB
}

func (p *PostgresConn) Close() error {
	return p.db.Close()
}

func NewPostgresConn(dataSource string) (*PostgresConn, error) {
	d, err := sql.Open("postgres", dataSource)
	if err == nil {
		err = d.Ping()
	}

	return &PostgresConn{db: d}, err
}
