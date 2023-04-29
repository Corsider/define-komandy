package structs

type GlobalTag struct {
	GlobalTagID int    `db:"globaltag_id"`
	Category    string `db:"category"`
}
