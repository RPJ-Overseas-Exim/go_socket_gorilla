package handlers

import (
	"RPJ_Overseas_Exim/go_mod_home/handlers/middlewares"
	"RPJ_Overseas_Exim/go_mod_home/services/socket"
	"RPJ_Overseas_Exim/go_mod_home/utils"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, mh *MessageHandler, ch *ChatHandler, ah *AuthHandler, adh *AdminHandler){
    // api routes

    e.GET("/", func (c echo.Context) error{
        return ch.renderMessages(c)
    })

    // socket routes

    hub := socket.NewHub()
    go hub.Run()

    e.GET("/ws", func (c echo.Context) error {
        email := c.QueryParam("email")

        if email==""{
            he := utils.HTTPException{Message: "Chat Id not given"}
            return &he
        }

        chatId,userId  := ch.cs.GetChatAndUserId(email)
        // log.Println("email: ", email, "ChatId: ", chatId, "UserId: ", userId)
        socket.ServeWs(chatId, userId, hub, c)

        return nil
    })

    // authentication routes

    e.GET("/login", ah.LoginView)
    e.POST("/login", ah.LoginHandler)
    e.POST("/logout", ah.LogoutHandler)

    // admin routes

    e.Use(middlewares.AuthUser)
    e.GET("/admin", adh.AdminView)
}
