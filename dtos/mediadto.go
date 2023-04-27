package dtos

import "github.com/chentihe/gin-mongo-api/models"

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

func ToModel(dto interface{}) *models.Media {
	var media *models.Media

	switch dto.(type) {
	case UpdateMediaDto:
		updateMediaDto := dto.(UpdateMediaDto)
		if updateMediaDto.Name != "" {
			media.Name = updateMediaDto.Name
		}
		if updateMediaDto.Thumbnail != "" {
			media.Thumbnail = updateMediaDto.Thumbnail
		}
		if updateMediaDto.Homepage != "" {
			media.Homepage = updateMediaDto.Homepage
		}
	case CreateMediaDto:
		createMediaDto := dto.(CreateMediaDto)
		media.Name = createMediaDto.Name
		media.Thumbnail = createMediaDto.Thumbnail
		media.Homepage = createMediaDto.Homepage
	}

	return media
}
