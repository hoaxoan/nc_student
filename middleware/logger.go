package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			log.Println(c.Request())
			return next(c)
		}
	}
}
