package daos

import (
	"context"
	"errors"
	"log"

	"github.com/chentihe/mongodb-api/models"
	"github.com/chentihe/mongodb-api/types"
	"github.com/chentihe/mongodb-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MediaCollectionName = `media`
)

type MediaDaoImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewMediaDao(db *mongo.Database, ctx context.Context) MediaDao {
	return &MediaDaoImpl{
		collection: db.Collection(MediaCollectionName),
		ctx:        ctx,
	}
}

func (m *MediaDaoImpl) GetAllMedia(paginate *types.MongoPaginate) (res *models.Medium, err error) {
	query := bson.M{}

	curr, err := m.collection.Find(m.ctx, query, utils.ToOpts(paginate))
	if err != nil {
		return nil, err
	}

	res = &models.Medium{
		Medium:     make([]*models.Media, 0, paginate.Limit),
		Pagination: paginate,
	}

	for curr.Next(m.ctx) {
		var media *models.Media
		if err := curr.Decode(&media); err != nil {
			log.Fatalln(err)
		}

		res.Medium = append(res.Medium, media)
	}

	return res, nil
}

func (m *MediaDaoImpl) GetMediaById(id primitive.ObjectID) (res *models.Media, err error) {
	query := bson.M{"_id": id}
	if err = m.collection.FindOne(m.ctx, query).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MediaDaoImpl) CreateMedia(media *models.Media) (res *models.Media, err error) {
	result, err := m.collection.InsertOne(m.ctx, media)
	if err != nil {
		return nil, err
	}

	query := bson.M{"_id": result.InsertedID}
	if err = m.collection.FindOne(m.ctx, query).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MediaDaoImpl) UpdateMediaById(id primitive.ObjectID, media *models.Media) (res *models.Media, err error) {
	docs, err := utils.ToDoc(media)
	opts := options.FindOneAndUpdate().SetReturnDocument(1)
	query := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: docs}}
	result := m.collection.FindOneAndUpdate(m.ctx, query, update, opts)
	if err = result.Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MediaDaoImpl) DeleteMediaById(id primitive.ObjectID) (err error) {
	query := bson.M{"_id": id}

	result, err := m.collection.DeleteOne(m.ctx, query)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
