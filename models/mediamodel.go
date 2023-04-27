package models

import (
	"time"

	"github.com/chentihe/gin-mongo-api/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Media struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Thumbnail string             `json:"thumbnail" bson:"thumbnail"`
	Homepage  string             `json:"homepage" bson:"homepage"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

type Medium struct {
	Medium     []*Media             `json:"medium"`
	Pagination *types.MongoPaginate `json:"pagination"`
}
