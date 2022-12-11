package routes

import (
	v1 "regin/routes/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"
)

func Init(r *gin.Engine, ctn *di.Container) {
	g := r.Group("/api/v1")
	{
		g.GET("dummy", v1.GetV1Response(ctn))
	}
}
