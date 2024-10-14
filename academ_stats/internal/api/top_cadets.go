package api

import (
	"academ_stats/internal/domain/response"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) TopCadets(c echo.Context) error {
	type module struct {
		ID int `query:"id"`
	}

	var m module
	if err := c.Bind(&m); err != nil {
		c.Set(logErr, fmt.Sprintf("TopCadets: query: \"%s\": %s", c.QueryString(), err))
		return customErrResponse(c, &response.ErrBadReq{Message: err.Error()}, nil)
	}

	// validate
	if m.ID <= 0 {
		c.Set(logErr, fmt.Sprintf("TopCadets: validate id: %d", m.ID))
		return customErrResponse(c, &response.ErrBadReq{Message: "bad id"}, nil)
	}

	// get dashboard
	top, err := h.svc.TopCadets(c.Request().Context(), m.ID)
	if err != nil {
		c.Set(logErr, fmt.Sprintf("TopCadets: %s", err))
		return customErrResponse(c, err, top)
	}

	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	return ok(c, top)
}

func (h *handlers) TopCadetsFile(c echo.Context) error {
	type module struct {
		ID int `query:"id"`
	}

	var m module
	if err := c.Bind(&m); err != nil {
		c.Set(logErr, fmt.Sprintf("TopCadetsFile: query: \"%s\": %s", c.QueryString(), err))
		return customErrResponse(c, &response.ErrBadReq{Message: err.Error()}, nil)
	}

	// validate
	if m.ID <= 0 {
		c.Set(logErr, fmt.Sprintf("TopCadetsFile: validate id: %d", m.ID))
		return customErrResponse(c, &response.ErrBadReq{Message: "bad id"}, nil)
	}

	// get dashboard file
	top, err := h.svc.TopCadetsFile(c.Request().Context(), m.ID)
	if err != nil {
		c.Set(logErr, fmt.Sprintf("TopCadets: %s", err))
		return customErrResponse(c, err, top)
	}

	c.Response().Header().Set("Content-Disposition", "attachment; filename=user_excel.xlsx")

	return c.Blob(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", top)
}
