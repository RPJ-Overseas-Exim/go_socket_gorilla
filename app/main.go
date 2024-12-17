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

    cs := services.NewChatService(db)
    ch := handlers.NewChatHandler(cs)

    handlers.SetupRoutes(e, mh, ch)
    e.Logger.Fatal(e.Start(":8181"))
}

    // e.GET("/ws", func (c echo.Context) error {
    //     email := c.QueryParam("email")
    //
    //     if email==""{
    //         he := utils.HTTPException{Message: "Chat Id not given"}
    //         return &he
    //     }
    //
    //     chatId := cs.GetChatId(email)
    //     socket.ServeWs(chatId, hub, c)
    //     return nil
    // })
