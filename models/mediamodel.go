package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MediaCollectionName = `media`
)

type (
	MediaModel interface {
		GetAllMedia(filter map[string]interface{}) ([]*Media, error)
		GetMediaByName(mediaName string) (*Media, error)
		CreateMedia(media *Media) (*Media, error)
		UpdateMediaById(id primitive.ObjectID, media *Media) (*Media, error)
		DeleteMediaById(id primitive.ObjectID) (*Media, error)
	}

	defaultMediaModel struct {
		collection *mongo.Collection
		ctx        context.Context
	}

	Media struct {
		Id        primitive.ObjectID `json:"id,omitempty"`
		Name      string             `json:"name,omitempty" validate:"required"`
		Thumbnail string             `json:"thumbnail,omitempty" validate:"required"`
		Homepage  string             `json:"homepage,omitempty" validate:"required"`
	}
)

func NewMediaModel(db *mongo.Database, ctx context.Context) MediaModel {
	return &defaultMediaModel{
		collection: db.Collection(MediaCollectionName),
		ctx:        ctx,
	}
}

func (m *defaultMediaModel) GetAllMedia(filter map[string]interface{}) (res []*Media, err error) {
	results, err := m.collection.Find(m.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for results.Next(m.ctx) {
		var media *Media
		if err = results.Decode(&media); err != nil {
			return nil, err
		}

		res = append(res, media)
	}

	return res, nil
}

func (m *defaultMediaModel) GetMediaByName(mediaName string) (res *Media, err error) {
	filter := bson.D{{Key: "name", Value: mediaName}}
	err = m.collection.FindOne(m.ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *defaultMediaModel) CreateMedia(media *Media) (res *Media, err error) {
	result, err := m.collection.InsertOne(m.ctx, media)
	if err != nil {
		return nil, err
	}

	res.Id = result.InsertedID.(primitive.ObjectID)

	return res, nil
}

func (m *defaultMediaModel) UpdateMediaById(id primitive.ObjectID, media *Media) (res *Media, err error) {
	_, err = m.collection.UpdateByID(m.ctx, id, media)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *defaultMediaModel) DeleteMediaById(id primitive.ObjectID) (res *Media, err error) {
	filter := bson.M{"id": id}

	_, err = m.collection.DeleteOne(m.ctx, filter)
	if err != nil {
		return nil, err
	}

	return res, nil
}
