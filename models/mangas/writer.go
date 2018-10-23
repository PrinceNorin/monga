package mangas

import (
	"github.com/PrinceNorin/monga/models"
	"github.com/PrinceNorin/monga/utils"
	"github.com/PrinceNorin/monga/utils/validators"
	"github.com/gin-gonic/gin"
)

func Create(f *validators.MangaForm, c *gin.Context) (*models.Manga, error) {
	var manga models.Manga
	if err := assignAttributes(&manga, f, c); err != nil {
		return nil, err
	}

	if err := models.ORM.Create(&manga).Error; err != nil {
		return nil, err
	}
	return &manga, nil
}

func Update(manga *models.Manga, f *validators.MangaForm, c *gin.Context) error {
	if err := assignAttributes(manga, f, c); err != nil {
		return err
	}

	if err := models.ORM.Save(manga).Error; err != nil {
		return err
	}
	return nil
}

func assignAttributes(manga *models.Manga, f *validators.MangaForm, c *gin.Context) error {
	manga.Title = f.Title
	manga.Desc = f.Desc
	manga.PublishedAt = f.PublishedAt
	manga.Status = f.Status

	cover, err := c.FormFile("cover")
	if err == nil && cover != nil {
		path, err := utils.UploadFile(cover, c)
		if err != nil {
			return err
		}
		manga.Cover = path
	}

	wallpaper, err := c.FormFile("wallpaper")
	if err == nil && wallpaper != nil {
		path, err := utils.UploadFile(wallpaper, c)
		if err != nil {
			return err
		}
		manga.Wallpaper = path
	}

	return nil
}
