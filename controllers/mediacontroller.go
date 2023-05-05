package controllers

import (
	"net/http"

	"github.com/chentihe/mongodb-api/config/svc"
	"github.com/chentihe/mongodb-api/dtos"
	"github.com/chentihe/mongodb-api/services"
	"github.com/chentihe/mongodb-api/types"
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
//
//	@Summary		Show All Media
//	@Description	get medium for the given page and limit
//	@Tags			Media
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			page	query		int	true	"Page"
//	@Param			limit	query		int	true	"Limit"
//	@Success		200		{object}	models.Medium
//	@Router			/media [get]
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

// GetAllMediaById godoc
//
//	@Summary		Show The Media By Id
//	@Description	get the media for the given id
//	@Tags			Media
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"Id"
//	@Success		200	{object}	models.Media
//	@Router			/media/{id} [get]
func (c *MediaController) GetMediaById(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.MediaService.GetMediaById(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusOK, res)
}

// CreateMedia godoc
//
//	@Summary		Create New Media
//	@Description	create the media
//	@Tags			Media
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		dtos.CreateMediaDto	true	"Create Media Request"
//	@Success		201		{object}	models.Media
//	@Router			/media [post]
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

// UpdateMedia godoc
//
//	@Summary		Update The Media By Id
//	@Description	update the media by id
//	@Tags			Media
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		string				true	"Id"
//	@Param			request	body		dtos.UpdateMediaDto	true	"Update Media Request"
//	@Success		200		{object}	models.Media
//	@Router			/media/{id} [put]
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

// DeleteMedia godoc
//
//	@Summary		Delete The Media By Id
//	@Description	delete the media by id
//	@Tags			Media
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	string	true	"Id"
//	@Success		203
//	@Router			/media/{id} [delete]
func (c *MediaController) DeleteMediaById(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.MediaService.DeleteMediaById(id); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, nil)
}
