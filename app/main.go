package main

import (
	"RPJ_Overseas_Exim/go_mod_home/db"
	"RPJ_Overseas_Exim/go_mod_home/handlers"
	"RPJ_Overseas_Exim/go_mod_home/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main(){
    e := echo.New()
    e.Use(middleware.Logger())
    e.Static("/static", "static")


    db := db.InitializeDB()
    ms := services.NewMessageService(db)
    mh := handlers.NewMessageHandler(ms)

    e.GET("/ws", func (c echo.Context) error{
        serveWs(newHub(), c)
        return nil
    })

    handlers.SetupRoutes(e, mh)
    e.Logger.Fatal(e.Start(":8181"))
}
