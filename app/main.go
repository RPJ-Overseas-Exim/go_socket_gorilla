package main

import (
	"RPJ_Overseas_Exim/go_mod_home/db"
	"RPJ_Overseas_Exim/go_mod_home/handlers"
	"RPJ_Overseas_Exim/go_mod_home/handlers/middlewares"
	"RPJ_Overseas_Exim/go_mod_home/services"
	"RPJ_Overseas_Exim/go_mod_home/services/socket"

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

    mid := middlewares.NewMiddleware(db)

    as := services.NewAuthService()
    ah := handlers.NewAuthHandler(as)

    cs := services.NewChatService(db)
    ch := handlers.NewChatHandler(cs)

    ads := services.NewAdminService()
    adh := handlers.NewAdminHandler(ads, cs)

    hub := socket.NewHub(db)
    handlers.SetupRoutes(e, hub, mh, ch, ah, adh, mid)
    e.Logger.Fatal(e.Start(":8181"))
}
