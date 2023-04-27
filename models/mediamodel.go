package models

import (
	"time"

	"github.com/chentihe/gin-mongo-api/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Media struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Thumbnail string             `bson:"thumbnail"`
	Homepage  string             `bson:"homepage"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

type Medium struct {
	Medium     []*Media
	Pagination *types.MongoPaginate
}
