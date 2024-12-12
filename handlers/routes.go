package handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo, mh *MessageHandler){

    e.GET("/", func (c echo.Context) error{
        return mh.renderLiveChat(c)
    })
}
