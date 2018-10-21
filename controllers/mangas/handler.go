package mangasController

import (
	"net/http"

	"github.com/PrinceNorin/monga/models/mangas"
	"github.com/PrinceNorin/monga/utils"
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

func ShowHandler(c *gin.Context) {
	id := utils.GetIntParam("mangaId", c)
	manga, err := mangas.Find(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, manga)
}
