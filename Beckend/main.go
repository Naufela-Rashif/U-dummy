package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		AppName: "U-dummy",
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
