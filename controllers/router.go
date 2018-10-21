package controllers

import (
	_ "github.com/PrinceNorin/monga/controllers/mangas"
	"github.com/PrinceNorin/monga/controllers/router"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = router.Get()
}
