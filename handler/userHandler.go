package handler

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/eminsit/openhr/config"
	"github.com/eminsit/openhr/db"
	"github.com/eminsit/openhr/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var ctx context.Context

var clt *mongo.Client

var database *mongo.Database

type UserHandler interface {
	create(user *model.User) model.User
	update(user *model.User) model.User
}

type JWTClaims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

func getDatabase() {
	clt = db.GetClient()
	database = clt.Database(config.AppConfig.DatabaseNames["auth"])
}

func Login(userLogin *model.UserLogin) (string, error) {
	getDatabase()
	userCollection := database.Collection(config.AppConfig.CollectionNames["user"])

	existedUser := model.User{}
	filter := bson.M{"username": userLogin.Username}

	err := userCollection.FindOne(ctx, filter).Decode(&existedUser)
	if err != nil {
		log.Println("Error while getting argument from DB", err)
	}
	if existedUser.Username == "" {
		return "", errors.New("Username Password does not match")
	}

	if !existedUser.ComparePasswords(userLogin.Password) {
		return "", errors.New("Username Password does not match")
	}

	jwtToken, err := generateJWTKey(existedUser)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func generateJWTKey(user model.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	user.PasswordEnc = ""
	claims := &JWTClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.JwtSecretKey))
	log.Println(config.AppConfig.JwtSecretKey)
	if err != nil {
		return "", errors.New("token generate error: " + err.Error())
	}

	return tokenString, nil
}

func CreateUser(user *model.User) (model.User, error) {
	getDatabase()
	userCollection := database.Collection(config.AppConfig.CollectionNames["user"])

	existedUser := model.User{}
	filter := bson.M{"username": user.Username}

	err := userCollection.FindOne(ctx, filter).Decode(&existedUser)
	if err != nil {
		log.Println("Error while getting argument from DB", err)
	}

	log.Printf("existedUser %v", existedUser)
	if existedUser.Username != "" {
		return model.User{}, errors.New("user already exists")
	}

	user.HashAndSaltPassword()
	user.Password = ""
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err.Error())
		return model.User{}, err
	}

	user.PasswordEnc = ""

	return *user, nil
}
