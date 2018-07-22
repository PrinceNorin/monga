package controllers

import (
	"github.com/PrinceNorin/monga/controllers/router"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = router.Get()
}
