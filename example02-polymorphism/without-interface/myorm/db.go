package myorm

import (
	"database/sql"
)

// MySQL is a struct that provides methods for performing database operations specific to MySQL.
type MySQL struct {
	Conn *sql.DB
}

// Insert is a method to insert data into mysql
func (m *MySQL) Insert(table string, data map[string]interface{}) error {
	// insert into mysql using mysql query
	return nil
}

// Postgres is a struct that provides methods for performing database operations specific to Postgres.
type Postgres struct {
	Conn *sql.DB
}

// Insert is a method to insert data into postgres
func (p *Postgres) Insert(table string, data map[string]interface{}) error {
	// insert into postgres using postgres query
	return nil
}
