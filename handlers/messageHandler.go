package handlers

import (
	views_livechat "RPJ_Overseas_Exim/go_mod_home/views/live_chat"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type MessageService interface {
    getMessages() []string;
    sendMessage(chatId, userId, content string) error;
}

type MessageHandler struct {
    MS *MessageService
}

func NewMessageHandler (ms *MessageService) *MessageHandler {
    return &MessageHandler{ms}
}

func (mh *MessageHandler) renderLiveChat (c echo.Context) error{
    comp := views_livechat.LiveChat()
    return renderView(c, 200, comp)
}

func renderView(c echo.Context, status int, comp templ.Component) error{
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
    c.Response().WriteHeader(status)

	return comp.Render(c.Request().Context(), c.Response().Writer)
}
