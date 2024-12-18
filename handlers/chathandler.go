package handlers

import (
	"RPJ_Overseas_Exim/go_mod_home/db"
	views_livechat "RPJ_Overseas_Exim/go_mod_home/views/live_chat"

	"github.com/labstack/echo/v4"
)

type chatService interface {
    GetChatAndUserId(string) (string,string);
    GetAllChats() *[]db.ResultsType;
}

type ChatHandler struct {
    cs chatService
}

func (ch *ChatHandler) renderMessages(c echo.Context)error{
    liveChat := views_livechat.LiveChat()
    return renderView(c, 200, liveChat)
}

func NewChatHandler(cs chatService)*ChatHandler{
    return &ChatHandler{
        cs,
    }
}
