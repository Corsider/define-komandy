package structs

import "github.com/lib/pq"

type User struct {
	UserId      int           `db:"user_id"`
	Name        string        `db:"name"`
	Nickname    string        `db:"nickname"`
	Rate        float64       `db:"rate"`
	Description string        `db:"description"`
	Friends     pq.Int32Array `db:"friends"`
	Logo        string        `db:"logo"`
	Media       string        `db:"media"`
	Tags        pq.Int32Array `db:"tags"`
	Mail        string        `db:"mail"`
	Password    string        `db:"password"`
	Region      string        `db:"region"`
}
