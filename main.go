package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"os"
	"web_uas/initializers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	engine := html.New("./views", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")
	app.Static("/css", "./public/assets/css")
	app.Static("/images", "./images")

	SetupRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))

	fmt.Print("Run on: localhost:3000")
}
