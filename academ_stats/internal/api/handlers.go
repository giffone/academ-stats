package api

import (
	"errors"
	"log"

	"academ_stats/internal/domain/response"
	"academ_stats/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var logErr = "logErr"

type Handlers interface {
	TopCadets(c echo.Context) error
	TopCadetsFile(c echo.Context) error
	GetTokenExpDate(c echo.Context) error
	ModuleList(c echo.Context) error
}

type handlers struct {
	svc  service.Service
	rLog bool
}

func NewHandlers(logg echo.Logger, svc service.Service, debug bool) Handlers {
	if debug {
		log.Println("request logging is enabled")
	}
	return &handlers{
		svc:  svc,
		rLog: debug,
	}
}

func customErrResponse(c echo.Context, err error, data any) error {
	defer printLogErr(c)

	if data == nil {
		data = []string{} // to show empty array
	}
	if errors.Is(err, response.ErrAccessDenied) {
		return c.JSON(http.StatusUnauthorized, response.Data{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
			Data:    data,
		})
	}
	var errBadReq *response.ErrBadReq
	if errors.As(err, &errBadReq) {
		return c.JSON(http.StatusBadRequest, response.Data{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    data,
		})
	}

	return c.JSON(http.StatusInternalServerError, response.Data{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
		Data:    data,
	})
}

func printLogErr(c echo.Context) {
	if eLog := c.Get(logErr); eLog != nil {
		log.Printf("\n//----\n[error]: %v\n----\\\\\n", eLog)
	}
}

func created(c echo.Context) error {
	return c.JSON(http.StatusCreated, response.Data{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    []string{},
	})
}

func ok(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, response.Data{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func noContent(c echo.Context, data any) error {
	if data == nil {
		data = []string{}
	}
	return c.JSON(http.StatusOK, response.Data{
		Code:    http.StatusNoContent,
		Message: "No content",
		Data:    data,
	})
}
