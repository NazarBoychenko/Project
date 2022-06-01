package event

type Event struct {
	EventId          int64  `db:"event_id,omitempty"`
	Title            string `db:"title"`
	ShortDescription string `db:"shortDescription"`
	Description      string `db:"description"`
	LongLat          string `db:"long_lat"`
	Images           string `db:"images"`
	Preview          string `db:"preview"`
}
