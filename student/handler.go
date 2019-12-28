package student

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type StudentHandler struct {
	StudentSrv *StudentService
}

func NewStudentHandler(e *echo.Echo, srv *StudentService) {
	handler := &StudentHandler{
		StudentSrv: srv,
	}
	g := e.Group("/api/student/v1/public")
	g.GET("/user/login", handler.Login)
	// e.POST("/v1/public/user/register", handler.Store)
	// e.PUT("/v1/private/user", handler.GetByID)
	// e.PUT("/v1/private/user", handler.Delete)
}

func (h *StudentHandler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
