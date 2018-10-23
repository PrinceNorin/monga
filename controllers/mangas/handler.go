package mangasController

import (
	"net/http"

	"github.com/PrinceNorin/monga/models/mangas"
	"github.com/PrinceNorin/monga/utils"
	"github.com/PrinceNorin/monga/utils/messages"
	"github.com/PrinceNorin/monga/utils/validators"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	orderBy := utils.GetOrderParam(c)
	page, limit := utils.GetPageParam(c)

	mangas, err := mangas.FindAll(page, limit, orderBy)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, mangas)
}

func CreateHandler(c *gin.Context) {
	var f validators.MangaForm
	validators.Bind(&f, c)

	msg := messages.GetMessages(c)
	if err := validators.Validate(&f, msg); err != nil {
		c.Error(msg)
		return
	}

	manga, err := mangas.Create(&f, c)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": manga,
	})
}

func UpdateHandler(c *gin.Context) {
	id := utils.GetIntParam("mangaId", c)
	manga, err := mangas.Find(id)
	if err != nil {
		c.Error(err)
		return
	}

	var f validators.MangaForm
	validators.Bind(&f, c)

	msg := messages.GetMessages(c)
	if err := validators.Validate(&f, msg); err != nil {
		c.Error(msg)
		return
	}

	if err := mangas.Update(manga, &f, c); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": manga,
	})
}

func ShowHandler(c *gin.Context) {
	id := utils.GetIntParam("mangaId", c)
	manga, err := mangas.Find(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, manga)
}
