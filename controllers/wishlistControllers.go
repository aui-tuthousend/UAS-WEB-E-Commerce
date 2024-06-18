package controllers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"web_uas/initializers"
	"web_uas/models"
)

func InsertIntoWishlist(c *fiber.Ctx) error {
	idWishlist := c.Params("idUser")
	idW, err := strconv.ParseUint(idWishlist, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	idBarang := c.Params("idProduct")
	idB, err := strconv.ParseUint(idBarang, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}

	var product models.Product
	if err := initializers.GetDB().First(&product, uint(idB)).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Product not found")
	}

	newDetailWislis := models.DetailWishlist{
		IdWishlist:   uint(idW),
		IdProduct:    product.ID,
		ProductName:  product.ProductName,
		ProductPrice: product.ProductPrice,
		Quantity:     1,
	}

	if err := initializers.GetDB().Create(&newDetailWislis).Error; err != nil {
		return err
	}

	return c.Redirect("/")
}
