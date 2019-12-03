package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/nc_student/db"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func TestDB(c echo.Context) error {
	db.Test()
	return c.String(http.StatusOK, "Testdb")
}
