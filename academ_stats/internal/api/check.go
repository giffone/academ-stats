package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (h *handlers) GetTokenExpDate(c echo.Context) error {
	tokenExpDate, err := h.svc.GetTokenExpDate(c.Request().Context())
	if err != nil {
		c.Set(logErr, fmt.Sprintf("GetTokenExpDate: %s", err))
		return customErrResponse(c, err, nil)
	}

	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	return ok(c, tokenExpDate)
}
