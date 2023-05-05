package routes

import (
	"github.com/chentihe/mongodb-api/config/svc"
	"github.com/chentihe/mongodb-api/controllers"
	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(v1 *gin.RouterGroup, ctx *svc.ServiceContext) {
	authGroup := v1.Group("/login")
	authController := controllers.NewAuthController(ctx)
	authGroup.POST("", authController.Login)
}
