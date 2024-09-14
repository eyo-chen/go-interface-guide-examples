package myorm

// DB is an interface that wraps the DB operations
type DB interface {
	Insert(table string, data map[string]interface{}) error
}

// ORM is a struct with a driver field of type DB interface
type ORM struct {
	db DB
}

// New initializes a new ORM with the given database driver.
func New(db DB) *ORM {
	return &ORM{db: db}
}

// Insert inserts data into the specified table using the configured database driver.
func (o *ORM) Insert(table string, data map[string]interface{}) error {
	return o.db.Insert(table, data)
}
