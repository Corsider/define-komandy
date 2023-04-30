package structs

import "github.com/lib/pq"

type User struct {
	UserId      int           `db:"user_id" json:"user_id"`
	Name        string        `db:"name" json:"name"`
	Nickname    string        `db:"nickname" json:"nickname"`
	Rate        float64       `db:"rate" json:"rate"`
	Description string        `db:"description" json:"description"`
	Friends     pq.Int32Array `db:"friends" json:"friends"`
	Logo        string        `db:"logo" json:"logo"`
	Media       string        `db:"media" json:"media"`
	Tags        pq.Int32Array `db:"tags" json:"tags"`
	Mail        string        `db:"mail" json:"mail"`
	Password    string        `db:"password" json:"password"`
	RegionID    string        `db:"region_id" json:"region_id"`
}
