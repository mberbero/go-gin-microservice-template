package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type instance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg instance

func Connect(dbname string) error {
	var mongoURI = "mongodb://localhost:27017/" + dbname

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	db := client.Database(dbname)

	if err != nil {
		return err
	}

	mg = instance{
		Client: client,
		Db:     db,
	}

	return nil
}

func GetDB() instance {
	return mg
}
