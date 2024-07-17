package controllers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"web_uas/initializers"
	"web_uas/models"
)

func InsertIntoWishlist(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	idUser := sess.Get("IDuser").(uint)

	idBarang := c.Params("idProduct")
	idB, err := strconv.ParseUint(idBarang, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}

	var wislis []models.DetailWishlist
	if err := initializers.GetDB().Where("id_wishlist = ?", idUser).Find(&wislis).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve wishlist")
	}

	for _, wis := range wislis {
		if wis.IdProduct == uint(idB) {
			wis.Quantity += 1

			if err := initializers.GetDB().Save(&wis).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to update quantity")
			}

			return c.Redirect("/")
		}
	}

	newDetailWislis := models.DetailWishlist{
		IdWishlist: idUser,
		IdProduct:  uint(idB),
		Quantity:   1,
	}

	if err := initializers.GetDB().Create(&newDetailWislis).Error; err != nil {
		return err
	}

	referer := c.Get("Referer", "/")

	return c.Redirect(referer)
}

func ShowWishList(c *fiber.Ctx) error {

	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	idUser := sess.Get("IDuser").(uint)

	var wislis []models.DetailWishlist
	query := initializers.GetDB().Model(&models.DetailWishlist{}).Where("id_wishlist = ?", idUser)

	if err := query.Find(&wislis).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("No photos found with the given conditions")
	}

	w := 0

	type wish struct {
		IdWishlist   uint
		IdProduct    uint
		ProductImage string
		ProductName  string
		ProductPrice int
		ProductStock int
		Quantity     int
	}

	var wisl []wish

	for _, wi := range wislis {
		var prod models.Product
		if err := initializers.GetDB().First(&prod, wi.IdProduct).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Product not found")
		}

		newWish := wish{
			IdWishlist:   wi.ID,
			IdProduct:    wi.IdProduct,
			ProductImage: prod.ProductImageCover,
			ProductName:  prod.ProductName,
			ProductPrice: prod.ProductPrice,
			ProductStock: prod.ProductStock,
			Quantity:     wi.Quantity,
		}

		w += (newWish.Quantity * newWish.ProductPrice)

		wisl = append(wisl, newWish)
	}

	return c.Render("main/cart", fiber.Map{"wishlists": wisl, "w": w})
}

func DeleteWislis(c *fiber.Ctx) error {
	idW := c.Params("idW")

	id, err := strconv.ParseUint(idW, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}

	var wishlistItem models.DetailWishlist
	if err := initializers.GetDB().Where("ID = ?", uint(id)).First(&wishlistItem).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Wishlist item not found")
	}

	if err := initializers.GetDB().Delete(&wishlistItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update quantity")
	}

	referer := c.Get("Referer", "/")

	return c.Redirect(referer)
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

	referer := c.Get("Referer", "/")
	if quantityInt == 0 {
		var wishlistItem models.DetailWishlist
		if err := initializers.GetDB().Where("id_product = ?", uint(productIDUint)).First(&wishlistItem).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Wishlist item not found")
		}

		if err := initializers.GetDB().Delete(&wishlistItem).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update quantity")
		}
	} else {

		var wishlistItem models.DetailWishlist
		if err := initializers.GetDB().Where("id_product = ?", uint(productIDUint)).First(&wishlistItem).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Wishlist item not found")
		}

		wishlistItem.Quantity = quantityInt

		if err := initializers.GetDB().Save(&wishlistItem).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update quantity")
		}
	}

	return c.Redirect(referer)
}
