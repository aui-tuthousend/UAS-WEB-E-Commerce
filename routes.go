package main

import (
	"github.com/gofiber/fiber/v2"
	"web_uas/controllers"
	"web_uas/initializers"
	"web_uas/models"
)

func SetupRoutes(app *fiber.App) {

	m := func(c *fiber.Ctx) error {
		store := controllers.Store

		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		idU := sess.Get("IDuser")
		if idU == nil {
			return c.Redirect("/login")
		}

		return c.Next()
	}

	n := func(c *fiber.Ctx) error {
		store := controllers.Store

		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		idU := sess.Get("IDuser")
		if idU != nil {
			return c.Redirect("/home")
		}

		return c.Next()
	}
	//auth := app.Group("/", m)

	app.Get("/regis", func(c *fiber.Ctx) error { return c.Render("sesi/regis", fiber.Map{}) })
	app.Post("/register", controllers.Register)

	app.Get("/login", func(c *fiber.Ctx) error { return c.Render("sesi/login", fiber.Map{}) })
	app.Post("/loginVerif", controllers.Login)
	app.Get("/", n, controllers.Landing)

	app.Get("/logout", m, controllers.LogOut)

	app.Post("/insertWishlist/:idUser/:idProduct", m, controllers.InsertIntoWishlist)
	app.Post("/checkout", m, controllers.Checkout)
	app.Post("/updateWishlistQuantity/:idDQ", m, controllers.UpdateWishlistQ)

	//app.Get("/createCategory", m, func(c *fiber.Ctx) error { return c.Render("produk/createCategory", fiber.Map{}) })
	app.Post("/storeCategory", m, controllers.AddKategori)

	app.Post("/storeProduct", m, controllers.StoreProduct)
	app.Get("/createProduct", m, func(c *fiber.Ctx) error {
		var categories []models.Category
		if err := initializers.GetDB().Find(&categories).Error; err != nil {
			return err
		}

		return c.Render("produk/createProduct", fiber.Map{"categories": categories})
	})
	app.Get("/home", m, controllers.ShowProduct)
	app.Get("/produk/:id", controllers.ViewProduct)

	app.Get("/wishlist/:idUser", m, controllers.ShowWishList)
	app.Get("/history/:idUser", m, controllers.ShowTransaction)
}
