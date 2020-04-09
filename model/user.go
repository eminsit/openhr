package model

type User struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Name     string `json:"name" bson:"name"`
	LastName string `json:"last_name" bson:"last_name"`
	Password string `json:"password" bson:"password"`
}
