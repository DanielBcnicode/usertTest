package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// PostgresConn hold the postgres database connection 
type PostgresConn struct {
	db *sql.DB
}

// Close the connection
func (p *PostgresConn) Close() error {
	return p.db.Close()
}

// NewPostgresConn is the constructor and also open the connection
func NewPostgresConn(dataSource string) (*PostgresConn, error) {
	d, err := sql.Open("postgres", dataSource)
	if err == nil {
		err = d.Ping()
	}

	return &PostgresConn{db: d}, err
}
