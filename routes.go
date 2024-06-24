package main

import (
	"github.com/gofiber/fiber/v2"
	"web_uas/controllers"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/insertWishlist/:idUser/:idProduct", controllers.InsertIntoWishlist)
	app.Post("/storeProduct", controllers.StoreProduct)
	app.Post("/checkout", controllers.Checkout)
	app.Post("/updateWishlistQuantity/:idDQ", controllers.UpdateWishlistQ)

	app.Get("/createProduct", func(c *fiber.Ctx) error { return c.Render("produk/createProduct", fiber.Map{}) })
	app.Get("/", controllers.ShowProduct)
	app.Get("/produk/:id", controllers.ViewProduct)

	app.Get("/wishlist/:idUser", controllers.ShowWishList)
	app.Get("/history/:idUser", controllers.ShowTransaction)

}
