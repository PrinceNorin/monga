package mangasController

import "github.com/PrinceNorin/monga/controllers/router"

func init() {
	r := router.Get()
	g := r.Group("/api/mangas")
	{
		g.GET("", IndexHandler)
		g.POST("", CreateHandler)
		g.GET("/:mangaId", ShowHandler)
		g.PUT("/:mangaId", UpdateHandler)
		g.PATCH("/:mangaId", UpdateHandler)
	}
}
