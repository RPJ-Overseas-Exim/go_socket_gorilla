package handlers

import (
	"RPJ_Overseas_Exim/go_mod_home/views/live_chat/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthServices interface{
    VerifyUser(username, password string) bool
}

type AuthHandler struct{
    as AuthServices
}

func NewAuthHandler(as AuthServices) *AuthHandler {
    return &AuthHandler{
        as,
    }
}


func (ah *AuthHandler) LoginView(c echo.Context) error {
    loginView := auth.Login()
    return renderView(c, http.StatusOK, loginView)
}

func (ah *AuthHandler) LoginHandler(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    verified := ah.as.VerifyUser(username, password)

    if verified {
        loginView := auth.Login()
        return renderView(c, http.StatusBadRequest, loginView)
    }

    return c.Redirect(http.StatusMovedPermanently, "/admin")
}

func (ah *AuthHandler) LogoutHandler(c echo.Context) error {
    return nil
}
