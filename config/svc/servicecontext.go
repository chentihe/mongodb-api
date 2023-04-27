package svc

import (
	"context"

	"github.com/chentihe/gin-mongo-api/config"
	"github.com/chentihe/gin-mongo-api/config/database"
	"github.com/chentihe/gin-mongo-api/daos"
	"github.com/chentihe/gin-mongo-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	DB           *mongo.Database
	MediaService *services.MediaService
}

func NewServiceContext(config *config.Config) *ServiceContext {
	ctx := context.TODO()
	db := database.ConnectDB(ctx, &config.DataBase)

	mediaDao := daos.NewMediaDao(db, ctx)
	mediaService := services.NewMediaService(&mediaDao)

	return &ServiceContext{
		DB:           db,
		MediaService: mediaService,
	}
}
