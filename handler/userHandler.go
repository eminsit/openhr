package handler

import (
	"context"
	"github.com/eminsit/openhr/config"
	"github.com/eminsit/openhr/db"
	"github.com/eminsit/openhr/model"
	"log"
)

type UserHandler interface {
	create(user *model.User) model.User
	update(user *model.User) model.User
}

func CreateUser(user *model.User) model.User {
	client := db.Client
	database := client.Database(config.AppConfig.DatabaseNames["auth"])
	userCollection := database.Collection(config.AppConfig.CollectionNames["user"])
	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err.Error())
	}

	return *user
}
