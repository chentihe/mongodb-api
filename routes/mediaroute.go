package routes

import (
	"github.com/chentihe/mongodb-api/controllers"
	"github.com/gin-gonic/gin"
)

func AddMediaRoutes(v1 *gin.RouterGroup, mediaController *controllers.MediaController) {
	mediaGroup := v1.Group("/media")
	mediaGroup.GET("", mediaController.GetAllMedia)
	mediaGroup.GET(":id", mediaController.GetMediaById)
	mediaGroup.POST("", mediaController.CreateMedia)
	mediaGroup.PUT(":id", mediaController.UpdateMediaById)
	mediaGroup.DELETE(":id", mediaController.DeleteMediaById)
}
