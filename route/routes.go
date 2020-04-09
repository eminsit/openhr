package route

import (
	"github.com/eminsit/openhr/handler"
	"github.com/eminsit/openhr/model"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/deneme", func(c echo.Context) error {
		log.Println("deneme")
		return c.String(http.StatusOK, "naber")
	})

	e.POST("user/register", func(c echo.Context) error {

		user := &model.User{}
		if err := c.Bind(user); err != nil {
			log.Fatalln(err.Error())
		}
		userHandler := handler.CreateUser(user)
		return c.JSON(http.StatusOK, userHandler)
	})
}
