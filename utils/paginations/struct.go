package paginations

import "github.com/jinzhu/gorm"

type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	ShowSQL bool
}

type Pagination struct {
	Count    int         `json:"count"`
	Pages    int         `json:"pages"`
	Records  interface{} `json:"records"`
	Offset   int         `json:"offset"`
	Limit    int         `json:"limit"`
	Page     int         `json:"page"`
	PrevPage int         `json:"prevPage"`
	NextPage int         `json:"nextPage"`
}
