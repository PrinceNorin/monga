package models

import "time"

type Manga struct {
	Model

	Themes []Theme `json:"themes" gorm:"many2many:manga_themes"`

	Title       string     `json:"title" gorm:"not null"`
	Desc        string     `json:"desc"`
	Cover       string     `json:"cover"`
	Wallpaper   string     `json:"wallpaper"`
	Status      string     `json:"status"`
	PublishedAt *time.Time `json:"publishedAt"`
}
