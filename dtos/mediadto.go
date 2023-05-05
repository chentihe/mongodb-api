package dtos

import "github.com/chentihe/mongodb-api/models"

type CreateMediaDto struct {
	Name      string `json:"name" binding:"required"`
	Thumbnail string `json:"thumbnail" binding:"required,url"`
	Homepage  string `json:"homepage" binding:"required,url"`
}

type UpdateMediaDto struct {
	Name      string `json:"name,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Homepage  string `json:"homepage,omitempty"`
}

func (dto *UpdateMediaDto) FillEmptyField(media *models.Media) {
	if dto.Name == "" {
		dto.Name = media.Name
	}

	if dto.Thumbnail == "" {
		dto.Thumbnail = media.Thumbnail
	}

	if dto.Homepage == "" {
		dto.Homepage = media.Homepage
	}
}

func ToModel(dto interface{}) interface{} {
	var media *models.Media

	switch dto.(type) {
	case *UpdateMediaDto:
		updateMediaDto := dto.(*UpdateMediaDto)
		media = &models.Media{
			Name:      updateMediaDto.Name,
			Thumbnail: updateMediaDto.Thumbnail,
			Homepage:  updateMediaDto.Homepage,
		}
	case *CreateMediaDto:
		createMediaDto := dto.(*CreateMediaDto)
		media = &models.Media{
			Name:      createMediaDto.Name,
			Thumbnail: createMediaDto.Thumbnail,
			Homepage:  createMediaDto.Homepage,
		}
	}

	return media
}
