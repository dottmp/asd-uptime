package db

import "time"

const (
	StatusUp   = "up"
	StatusDown = "down"
)

type Site struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type SiteCheck struct {
	ID             string    `json:"id"`
	SiteID         string    `json:"site_id"`
	Status         string    `json:"status"`
	ResponseTimeMs int64     `json:"response_time_ms"`
	Error          string    `json:"error,omitempty"`
	CheckedAt      time.Time `json:"checked_at"`
}
