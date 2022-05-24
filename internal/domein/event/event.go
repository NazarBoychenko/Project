package event

import "github.com/upper/db/v4/adapter/postgresql"

var settings = postgresql.ConnectionURL{
	Database: `postgres`,
	Host:     `127.0.0.1`,
	User:     `postgres`,
	Password: `postgres`,
}

type Event struct {
	EventId          int64  `db:"event_id"`
	Title            string `db:"title"`
	ShortDescription string `db:"shortDescription"`
	Description      string `db:"description"`
	LongLat          string `db:"long_lat"`
	Images           string `db:"images"`
	Preview          string `db:"preview"`
}
