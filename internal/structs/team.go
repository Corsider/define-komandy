package structs

import "github.com/lib/pq"

type Team struct {
	TeamID      int           `db:"team_id" json:"team_id"`
	Name        string        `db:"name" json:"name"`
	Rate        float32       `db:"rate" json:"rate"`
	Description string        `db:"description" json:"description"`
	Rules       string        `db:"rules" json:"rules"`
	Logo        string        `db:"logo" json:"logo"`
	Media       string        `db:"media" json:"media"`
	RegDate     string        `db:"reg_date" json:"reg_date"`
	Place       string        `db:"place" json:"place"`
	Tags        pq.Int32Array `db:"tags" json:"tags"`
	RegionID    int           `db:"region_id" json:"region_id"`
}
