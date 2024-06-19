package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"web_uas/initializers"
	"web_uas/models"
)

func Checkout(c *fiber.Ctx) error {
	selectedItems := c.FormValue("selectedItems")
	//userID := c.Params("userID").(uint) // Assuming userID is set in context, adjust as necessary

	if selectedItems == "" {
		return c.Status(http.StatusBadRequest).SendString("No items selected for checkout")
	}

	itemIDs := []string{selectedItems}
	var selectedProducts []models.DetailWishlist
	for _, itemID := range itemIDs {
		id, err := strconv.Atoi(itemID)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid item ID")
		}

		var item models.DetailWishlist
		if err := initializers.GetDB().First(&item, id).Error; err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Failed to retrieve item")
		}

		selectedProducts = append(selectedProducts, item)

		fmt.Println(item.ProductName)
	}

	return c.Redirect("/")
}
