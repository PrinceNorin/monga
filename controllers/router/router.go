package router

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	once   sync.Once
	engine *gin.Engine
)

func Get() *gin.Engine {
	once.Do(func() {
		engine = gin.Default()
	})
	return engine
}
