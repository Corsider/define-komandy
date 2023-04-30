package structs

type GlobalTag struct {
	GlobalTagID int    `db:"globaltag_id" json:"globaltag_id"`
	Category    string `db:"category" json:"category"`
}
