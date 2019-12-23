package user_service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(e *echo.Echo, srv *UserService) {
	handler := &UserHandler{
		UserService: srv,
	}
	g := e.Group("/api/student/v1/public")
	g.GET("/user/login", handler.Login)
	// e.POST("/v1/public/user/register", handler.Store)
	// e.PUT("/v1/private/user", handler.GetByID)
	// e.PUT("/v1/private/user", handler.Delete)
}

func (a *UserHandler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
