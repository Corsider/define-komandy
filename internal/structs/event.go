package structs

import "github.com/lib/pq"

type Event struct {
	EventID     int           `db:"event_id" json:"event_id"`
	Name        string        `db:"name" json:"name"`
	Description string        `db:"description" json:"description"`
	Date        string        `db:"date" json:"date"`
	FormatID    int           `db:"format_id" json:"format_id"`
	MainTheme   string        `db:"main_theme" json:"main_theme"`
	Media       string        `db:"media" json:"media"`
	Place       string        `db:"place" json:"place"`
	Url         string        `db:"url" json:"url"`
	Tags        pq.Int32Array `db:"tags" json:"tags"`
	RegionID    int           `db:"region_id" json:"region_id"`
	CreatorID   int           `db:"creator_id" json:"creator_id"`
}
