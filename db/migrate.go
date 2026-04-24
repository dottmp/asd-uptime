package db

import (
	"database/sql"
)

// type SiteCheck struct {
// 	ID             string    `json:"id"`
// 	SiteID         string    `json:"site_id"`
// 	Status         string    `json:"status"`
// 	ResponseTimeMs int64     `json:"response_time_ms"`
// 	Error          string    `json:"error,omitempty"`
// 	CheckedAt      time.Time `json:"checked_at"`
// }

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS sites (
					id TETXT PRIMARY KEY,
					name TEXT NOT NULL,
					url TEXT NOT NULL,
					created_at DATETIME DEFAULT CURRENT_TIMESTAMP
			)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS site_checks (
				id TEXT PRIMARY KEY,
				site_id TEXT NOT NULL,
				status TEXT NOT NULL,
				response_time_ms INTEGER NOT NULL,
				error TEXT,
				checked_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			)
	`)

	return err
}
