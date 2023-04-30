package structs

type Format struct {
	FormatId int    `db:"format_id" json:"format_id"`
	Format   string `db:"format" json:"format"`
}
