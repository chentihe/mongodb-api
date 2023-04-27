package routes

import (
	"github.com/chentihe/gin-mongo-api/config/svc"
	"github.com/chentihe/gin-mongo-api/controllers"
	"github.com/gin-gonic/gin"
)

func AddMediaRoutes(v1 *gin.RouterGroup, ctx *svc.ServiceContext) {
	mediaGroup := v1.Group("/media")
	mediaController := controllers.NewMediaController(ctx)
	mediaGroup.GET("/", mediaController.GetAllMedia())
	mediaGroup.GET("/:name", mediaController.GetMediaByName())
	mediaGroup.POST("/", mediaController.CreateMedia())
	mediaGroup.PUT("/:id", mediaController.UpdateMediaById())
	mediaGroup.DELETE("/:id", mediaController.DeleteMediaById())
}
