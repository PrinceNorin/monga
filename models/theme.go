package models

type Theme struct {
	Model

	Mangas []Manga `json:"mangas" gorm:"many2many:manga_themes"`

	Name string `json:"name" gorm:"unique;not null"`
}
