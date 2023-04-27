package daos

import (
	"github.com/chentihe/gin-mongo-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaDao interface {
	GetAllMedia(page int, limit int) (*models.Medium, error)
	GetMediaById(id primitive.ObjectID) (*models.Media, error)
	GetMediaByName(mediaName string) (*models.Media, error)
	CreateMedia(media *models.Media) (*models.Media, error)
	UpdateMediaById(id primitive.ObjectID, media *models.Media) (*models.Media, error)
	DeleteMediaById(id primitive.ObjectID) error
}
