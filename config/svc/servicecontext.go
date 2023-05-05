package svc

import (
	"context"

	"github.com/chentihe/mongodb-api/config"
	"github.com/chentihe/mongodb-api/config/database"
	"github.com/chentihe/mongodb-api/daos"
	"github.com/chentihe/mongodb-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	DB           *mongo.Database
	MediaService *services.MediaService
	UserService  *services.UserService
	Config       *config.Config
}

func NewServiceContext(config *config.Config, ctx context.Context) (*ServiceContext, error) {
	db, err := database.ConnectDB(ctx, &config.DataBase)
	if err != nil {
		return nil, err
	}

	mediaDao := daos.NewMediaDao(db, ctx)
	mediaService := services.NewMediaService(&mediaDao)

	userService := services.NewUserService()

	return &ServiceContext{
		DB:           db,
		MediaService: mediaService,
		UserService:  userService,
		Config:       config,
	}, nil
}
