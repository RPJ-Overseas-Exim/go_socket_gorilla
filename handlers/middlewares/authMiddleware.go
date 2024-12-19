package middlewares

import (
	"RPJ_Overseas_Exim/go_mod_home/db/models"
	"RPJ_Overseas_Exim/go_mod_home/services/cookie"
	"RPJ_Overseas_Exim/go_mod_home/services/jwt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Middleware struct{
    dbConn *gorm.DB
}

func NewMiddleware(dbConn *gorm.DB) *Middleware {
    return &Middleware{
            dbConn: dbConn,
        }
}

func (m *Middleware) AuthUser(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        err := godotenv.Load()
        if err!=nil {
            log.Println("Failed to load the env")
            return c.Redirect(http.StatusMovedPermanently, "/login")
        }

        // get the token string from the cookie
        tokenString, err := c.Cookie("Authentication")
        if err!=nil {
            return c.Redirect(http.StatusMovedPermanently, "/login")
        }

        // verify the token and get the email from the token string
        decoded, err := jwt.VerifyToken([]byte(os.Getenv("SECRET_KEY")), tokenString.Value)
        if err!=nil {
            c.SetCookie(cookie.DeleteCookie("Authentication", ""))
            return c.Redirect(http.StatusMovedPermanently, "/login")
        }

        // get the user id from the database
        var admin models.SocketUser
        m.dbConn.Find(&admin, "email=?", decoded)
        c.Set("AdminId", admin.Id)

        return next(c)
    }
}

func (m *Middleware) AuthLogin(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        err := godotenv.Load()
        if err!=nil {
            log.Println("Failed to load the env")
        }

        // check if the cookie is present or not 
        // if present redirect to /admin else move to /login
        _, err = c.Cookie("Authentication")
        if err==nil {
            return c.Redirect(http.StatusMovedPermanently, "/admin")
        }

        return next(c)
    }
}
