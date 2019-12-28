package main

import (
	"fmt"
	"log"

	"github.com/hoaxoan/nc_student/config"
	db "github.com/hoaxoan/nc_student/db"
	md "github.com/hoaxoan/nc_student/middleware"
	ss "github.com/hoaxoan/nc_student/student"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(md.SingleLogger())
	client, err := db.Connection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// User
	ssRepo := &ss.StudentRepository{client}
	srv := &ss.StudentService{ssRepo}
	ss.NewStudentHandler(e, srv)

	log.Println(e.Start(":9090"))
}
