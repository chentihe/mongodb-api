package services

import (
	"time"

	"github.com/chentihe/mongodb-api/daos"
	"github.com/chentihe/mongodb-api/dtos"
	"github.com/chentihe/mongodb-api/models"
	"github.com/chentihe/mongodb-api/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaService struct {
	MediaDao daos.MediaDao
}

func NewMediaService(mediaDao *daos.MediaDao) *MediaService {
	return &MediaService{
		MediaDao: *mediaDao,
	}
}

func (s *MediaService) GetAllMedia(paginate *types.MongoPaginate) (res *models.Medium, err error) {
	return s.MediaDao.GetAllMedia(paginate)
}

func (s *MediaService) GetMediaById(id string) (res *models.Media, err error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	return s.MediaDao.GetMediaById(objId)
}

func (s *MediaService) CreateMedia(dto *dtos.CreateMediaDto) (res *models.Media, err error) {
	media := dtos.ToModel(dto).(*models.Media)
	media.CreatedAt = time.Now()
	media.UpdatedAt = media.CreatedAt
	return s.MediaDao.CreateMedia(media)
}

func (s *MediaService) UpdateMediaById(id string, dto *dtos.UpdateMediaDto) (res *models.Media, err error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	oldMedia, err := s.MediaDao.GetMediaById(objId)
	if err != nil {
		return nil, err
	}

	dto.FillEmptyField(oldMedia)

	media := dtos.ToModel(dto).(*models.Media)
	media.UpdatedAt = time.Now()

	return s.MediaDao.UpdateMediaById(objId, media)
}

func (s *MediaService) DeleteMediaById(id string) (err error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	return s.MediaDao.DeleteMediaById(objId)
}
