package main

import (
	"RPJ_Overseas_Exim/go_mod_home/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main(){
    e := echo.New()
    e.Use(middleware.Logger())
    e.GET("/", func (c echo.Context) error{
        return c.String(200, "ok")
    })
    db.InitializeDB()
    e.Logger.Fatal(e.Start(":8181"))
}
