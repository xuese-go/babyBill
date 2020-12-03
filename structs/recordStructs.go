package structs

type Record struct {
	Uuid    string  `json:"uuid" form:"uuid" gorm:"primary_key"`
	Dates   string  `json:"dates" form:"dates"`
	Money   float64 `json:"money" form:"money"`
	Remarks string  `json:"remarks" form:"remarks"`
}

//更改表名称
func (Record) TableName() string {
	return "record_table"
}
