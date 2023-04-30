package structs

type Region struct {
	RegionID   int    `db:"region_id" json:"region_id"`
	CountryID  int    `db:"country_id" json:"country_id"`
	RegionName string `db:"region_name" json:"region_name"`
}
