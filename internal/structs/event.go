package structs

import "github.com/lib/pq"

type Event struct {
	EventID     int           `db:"event_id"`
	Name        int           `db:"name"`
	Description string        `db:"description"`
	Date        string        `db:"date"`
	FormatID    int           `db:"format_id"`
	MainTheme   string        `db:"main_theme"`
	Media       string        `db:"media"`
	Place       string        `db:"place"`
	Url         string        `db:"url"`
	Tags        pq.Int32Array `db:"tags"`
	RegionID    int           `db:"region_id"`
	CreatorID   int           `db:"creator_id"`
}
