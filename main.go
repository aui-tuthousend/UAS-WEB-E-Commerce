package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"os"
	"web_uas/helpers"
	"web_uas/initializers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	engine := html.New("./views", ".tmpl")
	engine.AddFunc("mul", helpers.Mul)
	engine.AddFunc("plus", helpers.Plus)
	engine.AddFunc("min", helpers.Min)
	app := fiber.New(fiber.Config{
		Views:     engine,
		BodyLimit: 10 * 1024 * 1024,
	})

	app.Static("/", "./public")
	app.Static("/css", "./public/assets/css")
	app.Static("/images", "./images")

	SetupRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))

	fmt.Print("Run on: localhost:3000")
}
