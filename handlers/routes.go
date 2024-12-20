package handlers

import (
	"RPJ_Overseas_Exim/go_mod_home/handlers/middlewares"
	"RPJ_Overseas_Exim/go_mod_home/services/socket"
	"RPJ_Overseas_Exim/go_mod_home/utils"
	admin_views "RPJ_Overseas_Exim/go_mod_home/views/live_chat/admin"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, hub *socket.Hub, mh *MessageHandler, ch *ChatHandler, ah *AuthHandler, adh *AdminHandler, mid *middlewares.Middleware){
    // api routes

    e.GET("/", ch.renderMessages)

    // socket routes

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

    authRoutes := e.Group("", mid.AuthLogin)
    authRoutes.GET("/login", ah.loginView)
    authRoutes.POST("/login", ah.loginHandler)
    e.GET("/logout", ah.logoutHandler)

    // admin routes

    adminRoutes := e.Group("/admin", mid.AuthUser)
    adminRoutes.GET("", adh.adminView)

    adminHash :=  make(map[string]*socket.ChatParticipant)

    adminRoutes.GET("/ws",func(c echo.Context) error{
        adminId, ok  := c.Get("AdminId").(string)
        if !ok{
            c.Response().WriteHeader(400)
            return &utils.HTTPException{Message: "Could not find adminId on context"}
        }

        adminCP, err := socket.ServeAdminWs(adminId, hub, c)
        adminHash[adminId] = adminCP
        return err
    })

    adminRoutes.GET("/chat/:chatId", func(c echo.Context) error {
        adminId, ok := c.Get("AdminId").(string)
        if !ok {
            c.Response().WriteHeader(400)
            return &utils.HTTPException{Message: "Could not find adminId on context"}
        }

        chatId := c.Param("chatId")
        msgs := adh.ms.GetMessages(chatId)

        // log.Println("Chat before ", *adminHash[adminId])
        socket.SwitchChats(adminHash[adminId], chatId, hub)
        // log.Println("Chat after ", *adminHash[adminId])

        chatView := admin_views.Chat("Chat heading", msgs, adminId)
        return renderView(c, http.StatusOK, chatView)
    })
}
