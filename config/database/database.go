package database

import (
	"context"
	"fmt"

	"github.com/chentihe/mongodb-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context, database *config.DataBase) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s.mongodb.net/?retryWrites=true&w=majority",
		database.UserName,
		database.Password,
		database.Cluster,
	)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	db := client.Database(database.Name)
	if db.RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}); err != nil {
		return nil, err
	}

	return db, nil
}
