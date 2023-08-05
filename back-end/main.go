package main

import (
	"survey/Database"
	"github.com/gofiber/fiber/v2"
	"survey/Routers"
)

func main(){
	Database.Connect()
	app := fiber.New()

	Routers.Initalize(app)

	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}