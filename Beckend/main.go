package main

import (
	"Udummy/database"
	"Udummy/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New(fiber.Config{
		AppName: "Udummy",
	})
	route.Router(app)
	err := app.Listen(":3000")
	if err != nil {
		return
	}

}
