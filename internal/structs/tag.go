package structs

type Tag struct {
	TagID       int    `db:"tag_id" json:"tag_id"`
	Activity    string `db:"activity" json:"activity"`
	GlobalTagID int    `db:"globaltag_id" json:"globaltag_id"`
}
