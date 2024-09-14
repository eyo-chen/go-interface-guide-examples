package myorm

import "fmt"

// ORM is a struct that holds a database and provides a method to insert data into a database.
type ORM struct {
	db any
}

// New initializes a new ORM with the given database.
func New(db any) *ORM {
	return &ORM{db: db}
}

// Insert inserts data into the specified table using the configured database.
func (o *ORM) Insert(table string, data map[string]interface{}) error {
	switch d := o.db.(type) {
	case MySQL:
		return d.Insert(table, data)
	case Postgres:
		return d.Insert(table, data)
	default:
		return fmt.Errorf("unsupported database driver")
	}
}
