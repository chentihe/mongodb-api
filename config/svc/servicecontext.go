package svc

import (
	"context"

	"github.com/chentihe/mongodb-api/config"
	"github.com/chentihe/mongodb-api/config/database"
	"github.com/chentihe/mongodb-api/controllers"
	"github.com/chentihe/mongodb-api/daos"
	"github.com/chentihe/mongodb-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	DB              *mongo.Database
	AuthController  *controllers.AuthController
	MediaController *controllers.MediaController
}

func NewServiceContext(config *config.Config, ctx context.Context) (*ServiceContext, error) {
	db, err := database.ConnectDB(ctx, &config.DataBase)
	if err != nil {
		return nil, err
	}

	mediaDao := daos.NewMediaDao(db, ctx)
	mediaService := services.NewMediaService(&mediaDao)
	mediaController := controllers.NewMediaController(mediaService)

	userService := services.NewUserService()
	authController := controllers.NewAuthController(userService, &config.Jwt)

	return &ServiceContext{
		DB:              db,
		AuthController:  authController,
		MediaController: mediaController,
	}, nil
}
