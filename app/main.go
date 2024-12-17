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

    as := services.NewAuthService()
    ah := handlers.NewAuthHandler(as)

    ads := services.NewAdminService()
    adh := handlers.NewAdminHandler(ads)

    cs := services.NewChatService(db)
    ch := handlers.NewChatHandler(cs)

    handlers.SetupRoutes(e, mh, ch, ah, adh)
    e.Logger.Fatal(e.Start(":8181"))
}
