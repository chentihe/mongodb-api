package daos

import (
	"github.com/chentihe/mongodb-api/models"
	"github.com/chentihe/mongodb-api/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaDao interface {
	GetAllMedia(paginate *types.MongoPaginate) (*models.Medium, error)
	GetMediaById(id primitive.ObjectID) (*models.Media, error)
	CreateMedia(media *models.Media) (*models.Media, error)
	UpdateMediaById(id primitive.ObjectID, media *models.Media) (*models.Media, error)
	DeleteMediaById(id primitive.ObjectID) error
}
