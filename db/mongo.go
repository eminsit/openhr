package db

import (
	"context"
	"fmt"
	"github.com/eminsit/openhr/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoConnectionUri string

func Init() {

	MongoConnectionUri = fmt.Sprintf("mongodb://%v:%v", config.AppConfig.Mongo.Host, config.AppConfig.Mongo.Port)

	_, err := mongo.NewClient(options.Client().ApplyURI(MongoConnectionUri))

	if err != nil {
		log.Println(err)
	}
}

func GetClient() *mongo.Client {

	clientOptions := options.Client().ApplyURI(MongoConnectionUri)
	log.Println(MongoConnectionUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client
}



