package structs

type Tag struct {
	TagID       int    `db:"tag_id"`
	Activity    string `db:"activity"`
	GlobalTagID int    `db:"globaltag_id"`
}
