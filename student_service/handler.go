package student_service

import (
	"github.com/labstack/echo/v4"
)

type StudentHandler struct {
	SRV *StudentService
}

func NewStudentHandler(e *echo.Echo, srv *StudentService) {
	// handler := &StudentHandler{
	// 	SRV: srv,
	//}
	g := e.Group("/api/student/v1/public")
	//g.GET("/user/login", handler.Login)
	// e.POST("/v1/public/user/register", handler.Store)
	// e.PUT("/v1/private/user", handler.GetByID)
	// e.PUT("/v1/private/user", handler.Delete)
}

// func (a *UserHandler) Login(c echo.Context) error {
// 	return c.String(http.StatusOK, "OK")
// }
