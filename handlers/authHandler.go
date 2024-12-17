package handlers

import (
	"RPJ_Overseas_Exim/go_mod_home/services/jwt"
	"RPJ_Overseas_Exim/go_mod_home/views/live_chat/auth"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
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
    err := godotenv.Load()
    if err != nil {
        loginView := auth.Login()
        return renderView(c, http.StatusBadRequest, loginView)
    }

    username := c.FormValue("username")
    password := c.FormValue("password")

    verified := ah.as.VerifyUser(username, password)
    if !verified {
        loginView := auth.Login()
        return renderView(c, http.StatusBadRequest, loginView)
    }

    cookie := new(http.Cookie)
    cookie.Name = "Authentication"

    token := jwt.CreateToken([]byte(os.Getenv("SECRET_KEY")), username)
    cookie.Value = token

    cookie.Expires = time.Now().Add(24 * time.Hour)
    c.SetCookie(cookie)

    c.Response().Header().Set("HX-Redirect", "/admin")
    return c.Redirect(http.StatusOK, "/admin")
}

func (ah *AuthHandler) LogoutHandler(c echo.Context) error {
    return nil
}
