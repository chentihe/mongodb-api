package controllers

import (
	"net/http"

	"github.com/chentihe/gin-mongo-api/config/svc"
	"github.com/chentihe/gin-mongo-api/dtos"
	"github.com/chentihe/gin-mongo-api/services"
	"github.com/chentihe/gin-mongo-api/types"
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

// GetAllMedia godoc
// @Summary Show All Media
// @Description get medium for the given page and limit
// @Tags medium
// @Accept json
// @Produce json
// @Success 200 {object} models.Medium
// @Router /media [get]
func (c *MediaController) GetAllMedia(ctx *gin.Context) {
	var paginate *types.MongoPaginate
	if err := ctx.ShouldBindQuery(&paginate); err != nil {
		panic(err)
	}

	res, err := c.MediaService.GetAllMedia(paginate)
	if err != nil {
		panic(err)
	}

	ctx.IndentedJSON(http.StatusOK, res)
}

func (c *MediaController) GetMediaById(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.MediaService.GetMediaById(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusOK, res)
}

func (c *MediaController) CreateMedia(ctx *gin.Context) {
	var dto *dtos.CreateMediaDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.MediaService.CreateMedia(dto)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, res)
}

func (c *MediaController) UpdateMediaById(ctx *gin.Context) {
	id := ctx.Param("id")
	var dto *dtos.UpdateMediaDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.MediaService.UpdateMediaById(id, dto)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, res)
}

func (c *MediaController) DeleteMediaById(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.MediaService.DeleteMediaById(id); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, nil)
}
