package mangas

import (
	"github.com/PrinceNorin/monga/models"
	"github.com/PrinceNorin/monga/utils/paginations"
)

func FindAll(page, limit int, orderBy []string) (*paginations.Pagination, error) {
	var mangas []models.Manga
	p := &paginations.Param{
		DB:      models.ORM,
		Page:    page,
		Limit:   limit,
		OrderBy: orderBy,
	}
	return paginations.Pagging(p, &mangas)
}

func Find(id uint) (*models.Manga, error) {
	var manga models.Manga
	if err := models.ORM.Find(&manga, id).Error; err != nil {
		return nil, err
	}
	return &manga, nil
}
