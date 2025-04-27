package router

import (
	"xixi-movie/handler"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	ms, err := handler.SetupMoiveServer()
	if err != nil {
		panic("")
	}
	r.GET("/search", ms.HandlerSearch)
	r.GET("/qb-login", ms.HandlerLoginQb)
	r.GET("/qb-add", ms.HandlerAddQb)
}
