package route

import (
	"github.com/eminsit/openhr/controller"
	"net/http"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("user/login", controller.Register)

	e.POST("user/register", controller.Login)
}
