package handlers

import (
	admin_views "RPJ_Overseas_Exim/go_mod_home/views/live_chat/admin"

	"github.com/labstack/echo/v4"
)

type AdminServices interface{

}

type AdminHandler struct{
    as AdminServices
    cs chatService
    ms MessageService
}

func NewAdminHandler(as AdminServices, cs chatService, ms MessageService) *AdminHandler{
    return &AdminHandler{
        as,
        cs,
        ms,
    }
}

func (ah *AdminHandler) adminView(c echo.Context) error {
    users, seenMap := ah.cs.GetAllChats() 

    return renderView(c, 200, admin_views.AdminHome(users,seenMap))
}

