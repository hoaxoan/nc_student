package http

import (
	"net/http"

	"github.com/hoaxoan/nc_student/model"
	"github.com/hoaxoan/nc_student/student"
	"github.com/labstack/echo/v4"
)

type studentHandler struct {
	SUcase student.Usecase
}

func NewStudentHandler(e *echo.Echo, uc student.Usecase) {
	handler := &studentHandler{
		SUcase: uc,
	}
	PrivateRoute(e, handler)

	StaffRoute(e, handler)

	PublicRoute(e, handler)
}

func PrivateRoute(e *echo.Echo, handler *studentHandler) {
	//g := e.Group("/api/student/v1/private")
}

func StaffRoute(e *echo.Echo, handler *studentHandler) {
	g := e.Group("/api/student/v1/staff")
	g.POST("/student", handler.Create)
	g.PUT("/student", handler.Update)
	g.DELETE("/student", handler.Delete)
}

func PublicRoute(e *echo.Echo, handler *studentHandler) {
	g := e.Group("/api/student/v1/public")
	g.GET("/student", handler.GetAll)
	g.PATCH("/student", handler.Search)
	g.PATCH("/health", handler.Health)
}

func (h *studentHandler) Health(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}

func (h *studentHandler) GetAll(ctx echo.Context) error {
	var req model.StudentRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var res model.StudentResponse
	err := h.SUcase.GetAll(ctx.Request().Context(), &req, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, res.Students)
}

func (h *studentHandler) Search(ctx echo.Context) error {
	var req model.StudentRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var res model.StudentResponse
	err := h.SUcase.GetAll(ctx.Request().Context(), &req, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, res.Students)
}

func (h *studentHandler) Create(ctx echo.Context) error {
	var student model.Student
	if err := ctx.Bind(&student); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var res model.StudentResponse
	err := h.SUcase.Create(ctx.Request().Context(), &student, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, res.Student)
}

func (h *studentHandler) Update(ctx echo.Context) error {
	var student model.Student
	if err := ctx.Bind(&student); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var res model.StudentResponse
	err := h.SUcase.Update(ctx.Request().Context(), &student, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, res.Student)
}

func (h *studentHandler) Delete(ctx echo.Context) error {
	var student model.Student
	if err := ctx.Bind(&student); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var res model.StudentResponse
	err := h.SUcase.Delete(ctx.Request().Context(), &student, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, res.Student)
}
