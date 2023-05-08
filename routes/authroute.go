package routes

import (
	"github.com/chentihe/mongodb-api/controllers"
	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(v1 *gin.RouterGroup, authController *controllers.AuthController) {
	authGroup := v1.Group("/login")
	authGroup.POST("", authController.Login)
}
