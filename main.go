package main

import (
	"fmt"
	"github.com/eminsit/openhr/db"
	"reflect"

	"github.com/eminsit/openhr/config"
	"github.com/eminsit/openhr/route"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.Init()
	db.Init()
	route.Init(e)
	fmt.Println(reflect.TypeOf(e))
	e.Logger.Fatal(e.Start(":" + config.AppConfig.App.Port))
}
