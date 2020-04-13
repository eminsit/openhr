package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Username    string `json:"username" bson:"username"`
	Email       string `json:"email" bson:"email"`
	Name        string `json:"name" bson:"name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Password    string `json:"password,omitempty" bson:"-,omitempty"`
	PasswordEnc string `json:"password_encrypted,omitempty" bson:"password_encrypted"`
}

type UserLogin struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
}


func (u *User) HashAndSaltPassword() {
	pwd := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println("Encrypt password error: &v", err)
	}
	u.PasswordEnc = string(hash)
}

func (u *User) ComparePasswords(password string) bool {
	plainPwd := []byte(password)
	byteHash := []byte(u.PasswordEnc)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println("Comparing password error: %v", err)
		return false
	}

	return true
}

