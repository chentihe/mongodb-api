package dtos

import "github.com/chentihe/gin-mongo-api/models"

type CreateMediaDto struct {
	Name      string `form:"name"`
	Thumbnail string `form:"thumbnail"`
	Homepage  string `form:"homepage"`
}

func (dto *CreateMediaDto) ToModel() *models.Media {
	return &models.Media{
		Name:      dto.Name,
		Thumbnail: dto.Thumbnail,
		Homepage:  dto.Homepage,
	}
}

type UpdateMediaDto struct {
	Name      string `form:"name"`
	Thumbnail string `form:"thumbnail"`
	Homepage  string `form:"homepage"`
}

func (dto *UpdateMediaDto) ToModel() *models.Media {
	return &models.Media{
		Name:      dto.Name,
		Thumbnail: dto.Thumbnail,
		Homepage:  dto.Homepage,
	}
}
