package routes

import (
	"github.com/chentihe/gin-mongo-api/config/svc"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(router *gin.Engine, ctx *svc.ServiceContext) {
	v1 := router.Group("/v1")
	AddMediaRoutes(v1, ctx)
}
