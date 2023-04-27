package controllers

import (
	"net/http"

	"github.com/chentihe/gin-mongo-api/config/svc"
	"github.com/chentihe/gin-mongo-api/dtos"
	"github.com/chentihe/gin-mongo-api/services"
	"github.com/gin-gonic/gin"
)

type MediaController struct {
	MediaService *services.MediaService
}

func NewMediaController(svc *svc.ServiceContext) *MediaController {
	return &MediaController{
		MediaService: svc.MediaService,
	}
}

func (c *MediaController) GetAllMedia() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page := ctx.DefaultQuery("page", "1")
		limit := ctx.DefaultQuery("limit", "10")

		filter := make(map[string]interface{}, 0)
		filter["page"] = page
		filter["limit"] = limit

		res, err := c.MediaService.GetAllMedia(filter)
		if err != nil {
			panic(err)
		}

		ctx.IndentedJSON(http.StatusOK, res)
	}
}

func (c *MediaController) GetMediaByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mediaName := ctx.Query("mediaName")

		res, err := c.MediaService.GetMediaByName(mediaName)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		}

		ctx.IndentedJSON(http.StatusOK, res)
	}
}

func (c *MediaController) CreateMedia() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto dtos.CreateMediaDto

		if err := ctx.ShouldBindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newMedia := dto.ToModel()

		res, err := c.MediaService.CreateMedia(newMedia)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusCreated, res)
	}
}

func (c *MediaController) UpdateMediaById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var dto dtos.UpdateMediaDto

		if err := ctx.ShouldBindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateMedia := dto.ToModel()

		res, err := c.MediaService.UpdateMediaById(id, updateMedia)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, res)
	}
}

func (c *MediaController) DeleteMediaById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		res, err := c.MediaService.DeleteMediaById(id)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, res)
	}
}
