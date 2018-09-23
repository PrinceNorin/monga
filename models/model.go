package models

import "time"

type Model struct {
	ID        uint      `json:"id" gorm:"type:bigserial;primary_key"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp"`
}
