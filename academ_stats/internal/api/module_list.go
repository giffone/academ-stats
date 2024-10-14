package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (h *handlers) ModuleList(c echo.Context) error {
	// get dashboard
	list, err := h.svc.ModuleList(c.Request().Context())
	if err != nil {
		c.Set(logErr, fmt.Sprintf("ModuleList: %s", err))
		return customErrResponse(c, err, list)
	}

	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	return ok(c, list)
}