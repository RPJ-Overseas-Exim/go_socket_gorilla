package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminServices interface{
}

type AdminHandler struct{
    as AdminServices
}

func NewAdminHandler(as AdminServices) *AdminHandler{
    return &AdminHandler{
        as,
    }
}

func (ah *AdminHandler) AdminView(c echo.Context) error {
    return c.String(http.StatusOK, "Admin View route")
}
