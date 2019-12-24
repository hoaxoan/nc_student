package main

import (
	"fmt"
	"github.com/hoaxoan/nc_student/config"
	db "github.com/hoaxoan/nc_student/db"
	md "github.com/hoaxoan/nc_student/middleware"
	us "github.com/hoaxoan/nc_student/user_service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
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

	usRepo := &us.UserRepository{client}
	srv := &us.UserService{usRepo}
	us.NewUserHandler(e, srv)
	//route.All(e)

	log.Println(e.Start(":9090"))
}
