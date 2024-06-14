package main

import (
	"github.com/gofiber/fiber/v2"
	"web_uas/controllers"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/storeProduct", controllers.StoreProduct)

	app.Get("/createProduct", func(c *fiber.Ctx) error { return c.Render("produk/createProduct", fiber.Map{}) })
	app.Get("/", controllers.ShowProduct)
	app.Get("/produk/:id", controllers.ViewProduct)

}
