package services

import (
	"time"

	"github.com/chentihe/gin-mongo-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaService struct {
	MediaModel models.MediaModel
}

func NewMediaService(mediaModel *models.MediaModel) *MediaService {
	return &MediaService{
		MediaModel: *mediaModel,
	}
}

func (s *MediaService) GetAllMedia(filter map[string]interface{}) (res []*models.Media, err error) {
	return s.MediaModel.GetAllMedia(filter)
}

func (s *MediaService) GetMediaByName(mediaName string) (res *models.Media, err error) {
	return s.MediaModel.GetMediaByName(mediaName)
}

func (s *MediaService) CreateMedia(media *models.Media) (res *models.Media, err error) {
	media.CreatedAt = time.Now()
	media.UpdatedAt = media.CreatedAt
	return s.MediaModel.CreateMedia(media)
}

func (s *MediaService) UpdateMediaById(id string, media *models.Media) (res *models.Media, err error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	return s.MediaModel.UpdateMediaById(objId, media)
}

func (s *MediaService) DeleteMediaById(id string) (res *models.Media, err error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	return s.MediaModel.DeleteMediaById(objId)
}
