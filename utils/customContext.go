package utils

import "github.com/labstack/echo/v4"

type CustomContext struct{
    echo.Context

    UserId string
}

