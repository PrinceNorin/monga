package validators

import "time"

type MangaForm struct {
	Title       string     `form:"title" json:"title" validate:"required"`
	Desc        string     `form:"desc" json:"desc"`
	Status      string     `form:"status" json:"status" validate:"required,oneof=ongoing finished"`
	PublishedAt *time.Time `form:"publishedAt" json:"publishedAt" validate:"omitempty,date"`
}
