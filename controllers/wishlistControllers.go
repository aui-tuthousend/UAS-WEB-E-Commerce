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
		ProductImage: product.ProductImageCover,
		ProductName:  product.ProductName,
		ProductPrice: product.ProductPrice,
		Quantity:     1,
	}

	if err := initializers.GetDB().Create(&newDetailWislis).Error; err != nil {
		return err
	}

	return c.Redirect("/")
}

func ShowWishList(c *fiber.Ctx) error {
	idWishlist := c.Params("idUser")
	idW, err := strconv.ParseUint(idWishlist, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	var wislis []models.DetailWishlist
	query := initializers.GetDB().Model(&models.DetailWishlist{}).Where("id_wishlist = ?", uint(idW))

	if err := query.Find(&wislis).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("No photos found with the given conditions")
	}

	return c.Render("main/wishList", fiber.Map{"wishlists": wislis})
}

func UpdateWishlistQ(c *fiber.Ctx) error {
	productId := c.FormValue("productId")
	quantity := c.FormValue("quantity")

	productIDUint, err := strconv.ParseUint(productId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}

	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid quantity value")
	}

	var wishlistItem models.DetailWishlist
	if err := initializers.GetDB().Where("id_product = ?", uint(productIDUint)).First(&wishlistItem).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Wishlist item not found")
	}

	wishlistItem.Quantity = quantityInt

	if err := initializers.GetDB().Save(&wishlistItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update quantity")
	}

	referer := c.Get("Referer", "/")
	return c.Redirect(referer)
}
