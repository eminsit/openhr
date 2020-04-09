package db

import (
	"context"
	"fmt"
	"github.com/eminsit/openhr/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Client *mongo.Client

func Init() {

	mongoConnection := fmt.Sprintf("mongodb://%v:%v", config.AppConfig.Mongo.Host, config.AppConfig.Mongo.Port)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoConnection))

	if err != nil {

	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	Client = client
}
