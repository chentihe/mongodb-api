package svc

import (
	"context"

	"github.com/chentihe/gin-mongo-api/config/database"
	"github.com/chentihe/gin-mongo-api/models"
	"github.com/chentihe/gin-mongo-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	DB           *mongo.Database
	MediaService *services.MediaService
}

func NewServiceContext() *ServiceContext {
	ctx := context.Background()
	db := database.ConnectDB(ctx)

	mediaModel := models.NewMediaModel(db, ctx)
	mediaService := services.NewMediaService(&mediaModel)

	return &ServiceContext{
		DB:           db,
		MediaService: mediaService,
	}
}
