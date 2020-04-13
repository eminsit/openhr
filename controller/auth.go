package controller

import (
	"github.com/eminsit/openhr/handler"
	"github.com/eminsit/openhr/model"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func Register(c echo.Context) error {
	userLogin := &model.UserLogin{}
	if err := c.Bind(userLogin); err != nil {
		log.Fatalln(err.Error())
	}

	jwtToken, err := handler.Login(userLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"key": jwtToken,})
}

func Login(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		log.Fatalln(err.Error())
	}

	userHandler, err := handler.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userHandler)
}