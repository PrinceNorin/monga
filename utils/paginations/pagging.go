package paginations

import (
	"math"

	"github.com/jinzhu/gorm"
)

func Pagging(p *Param, data interface{}) (*Pagination, error) {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 30
	}
	if len(p.OrderBy) > 0 {
		for _, order := range p.OrderBy {
			db = db.Order(order)
		}
	}

	done := make(chan bool, 1)
	var pagination Pagination
	var count int
	var offset int

	go countRecords(db, data, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	if err := db.Limit(p.Limit).Offset(offset).Find(data).Error; err != nil {
		<-done
		return nil, err
	}
	<-done

	pagination.Count = count
	pagination.Records = data
	pagination.Page = p.Page

	pagination.Offset = offset
	pagination.Limit = p.Limit
	pagination.Pages = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		pagination.PrevPage = p.Page - 1
	} else {
		pagination.PrevPage = p.Page
	}

	if p.Page >= pagination.Pages {
		pagination.NextPage = p.Page
	} else {
		pagination.NextPage = p.Page + 1
	}
	return &pagination, nil
}

func countRecords(db *gorm.DB, data interface{}, done chan bool, count *int) {
	db.Model(data).Count(count)
	done <- true
}
