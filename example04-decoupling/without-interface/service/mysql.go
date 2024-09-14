package service

import "database/sql"

// MySQLService is a service that wrap the basic method of the MySQL
type MySQLService struct {
	sql *sql.DB
}

// NewMySQLService initialize a new MySQLService
func NewMySQLService(sql *sql.DB) *MySQLService {
	return &MySQLService{sql: sql}
}

// GetUserName is a method that gets the user name from the database
func (s *MySQLService) GetUserName(id string) (string, error) {
	// get user name from database
	row := s.sql.QueryRow("SELECT name FROM users WHERE id = ?", id)
	var name string
	err := row.Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}
