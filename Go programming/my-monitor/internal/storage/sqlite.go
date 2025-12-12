package storage

import (
	"database/sql"
	"time"

	// Import the driver anonymously so it registers itself with database/sql
	_ "modernc.org/sqlite"
)

// Record represents a single row in our database
type Record struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	Status    string    `json:"status"`
	CheckedAt time.Time `json:"checked_at"`
}

// DB manages the database connection
type DB struct {
	sqlDB *sql.DB
}

// NewSQLite initializes the database and creates the table if it doesn't exist
func NewSQLite(filepath string) (*DB, error) {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	// Create the table if it's missing
	query := `
	CREATE TABLE IF NOT EXISTS checks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT,
		status TEXT,
		checked_at DATETIME
	);`
	
	if _, err := db.Exec(query); err != nil {
		return nil, err
	}

	return &DB{sqlDB: db}, nil
}

// Save inserts a new record
func (d *DB) Save(url, status string) error {
	query := `INSERT INTO checks (url, status, checked_at) VALUES (?, ?, ?)`
	_, err := d.sqlDB.Exec(query, url, status, time.Now())
	return err
}

// GetAll retrieves the last 100 checks
func (d *DB) GetAll() ([]Record, error) {
	query := `SELECT id, url, status, checked_at FROM checks ORDER BY id DESC LIMIT 100`
	rows, err := d.sqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var r Record
		// We must scan pointers to the fields
		if err := rows.Scan(&r.ID, &r.URL, &r.Status, &r.CheckedAt); err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	return records, nil
}