package main

import (
	"fmt"
	"log"

	"github.com/hoaxoan/nc_student/config"
	"github.com/hoaxoan/nc_student/db"
	md "github.com/hoaxoan/nc_student/middleware"
	_studentHttpDeliver "github.com/hoaxoan/nc_student/student/delivery/http"
	_studentRepository "github.com/hoaxoan/nc_student/student/repository"
	_studentUsecase "github.com/hoaxoan/nc_student/student/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(md.Logger())
	client, err := db.Connection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	ssRepo := _studentRepository.NewStudentRepository(client)
	su := _studentUsecase.NewStudentUsecase(ssRepo)
	_studentHttpDeliver.NewStudentHandler(e, su)

	log.Println(e.Start(":9090"))
}
