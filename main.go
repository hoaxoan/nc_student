package main

import (
	"fmt"
	"log"

	"github.com/hoaxoan/nc_student/config"
	db "github.com/hoaxoan/nc_student/db"
	us "github.com/hoaxoan/nc_student/user_service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)

	e := echo.New()
	e.Use(middleware.Recover())

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
