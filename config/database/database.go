package database

import (
	"context"
	"fmt"

	"github.com/chentihe/gin-mongo-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context, database *config.DataBase) *mongo.Database {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.anx106i.mongodb.net/?retryWrites=true&w=majority",
		database.UserName, database.Password)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	db := client.Database(database.Name)
	if db.RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}); err != nil {
		panic(err)
	}

	return db
}
