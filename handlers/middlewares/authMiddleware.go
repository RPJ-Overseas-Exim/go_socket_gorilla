package middlewares

import (
	"RPJ_Overseas_Exim/go_mod_home/services/jwt"
	"os"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func AuthUser(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        err := godotenv.Load()
        if err!=nil {
            log.Println("Failed to load the env")
            return c.Redirect(http.StatusMovedPermanently, "/login")
        }

        tokenString, err := c.Cookie("Authentication")
        if err!=nil {
            return c.Redirect(http.StatusMovedPermanently, "/login")
        }

        decoded, err := jwt.VerifyToken([]byte(os.Getenv("SECRET_KEY")), tokenString.Value)
        if err!=nil {
            c.Response().Header().Set("HX-Redirect", "/login")
            return c.Redirect(http.StatusOK, "/login")
        }

        log.Printf("Token: %v", decoded.Claims)
        log.Printf("Cookie %v", tokenString.Value)

        return next(c)
    }
}
