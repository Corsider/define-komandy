package structs

type User struct {
	UserId      int     `db:"user_id"`
	Name        string  `db:"name"`
	Nickname    string  `db:"nickname"`
	Rate        float64 `db:"rate"`
	Description string  `db:"description"`
	Friends     []int   `db:"friends"`
	Tags        []int   `db:"tags"`
	Mail        string  `db:"mail"`
	Password    string  `db:"password"`
	Salt        string  `db:"salt"`
}
