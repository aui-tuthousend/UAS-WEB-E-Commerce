package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"web_uas/initializers"
	"web_uas/models"
)

type CheckoutPayload struct {
	SelectedItems []string `json:"selectedItems"`
}

func Checkout(c *fiber.Ctx) error {
	var payload CheckoutPayload

	if err := c.BodyParser(&payload); err != nil {
		log.Println("Error parsing JSON:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	total := 0

	newTransaction := models.Transaction{
		IdUser: 1,
	}

	if err := initializers.GetDB().Create(&newTransaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create transaction")
	}

	for _, wislis := range payload.SelectedItems {
		id, err := strconv.ParseUint(wislis, 10, 32)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
		}

		var DWislis models.DetailWishlist
		if err := initializers.GetDB().First(&DWislis, uint(id)).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Product not found")
		}

		var productReal models.Product
		if err := initializers.GetDB().First(&productReal, DWislis.IdProduct).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Product not found")
		}

		temptot := DWislis.Quantity * DWislis.ProductPrice
		total += temptot

		newProductCopies := models.ProductCopy{
			ProductName:        productReal.ProductName,
			ProductDescription: productReal.ProductDescription,
			ProductImageCover:  productReal.ProductImageCover,
			Quantity:           DWislis.Quantity,
			ProductPrice:       productReal.ProductPrice,
		}

		if err := initializers.GetDB().Create(&newProductCopies).Error; err != nil {
			return err
		}

		newDT := models.DetailTransaction{
			IdTransaction:   newTransaction.ID,
			IdProductCopies: newProductCopies.ID,
			Quantity:        DWislis.Quantity,
			Total:           temptot,
		}

		if err := initializers.GetDB().Create(&newDT).Error; err != nil {
			return err
		}

		if err := initializers.GetDB().Delete(&models.DetailWishlist{}, DWislis.ID).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete product")
		}
	}

	newTransaction.TotalPrice = total

	if err := initializers.GetDB().Save(&newTransaction).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Checkout successful",
	})
}
