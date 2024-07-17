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

func ShowTransaction(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	idUser := sess.Get("IDuser").(uint)

	var transactions []models.Transaction
	query := initializers.GetDB().Model(&models.Transaction{}).Where("id_user = ?", idUser)
	if err := query.Find(&transactions).Error; err != nil {
		return c.SendString("tidak adada")
	}

	var products []models.ProductCopy
	if err := initializers.GetDB().Find(&products).Error; err != nil {
		return err
	}

	var ditelTrans []models.DetailTransaction
	if err := initializers.GetDB().Find(&ditelTrans).Error; err != nil {
		return err
	}

	type pivot struct {
		DT models.DetailTransaction
		PC models.ProductCopy
	}

	type cooked struct {
		Transaksi models.Transaction
		Ditel     []pivot
	}

	var served []cooked

	for _, tra := range transactions {
		newCook := cooked{
			Transaksi: tra,
		}

		for _, dit := range ditelTrans {
			if dit.IdTransaction == tra.ID {
				newPivot := pivot{
					DT: dit,
				}

				for _, pc := range products {
					if dit.IdProductCopies == pc.ID {
						newPivot.PC = pc
					}
				}

				newCook.Ditel = append(newCook.Ditel, newPivot)
			}
		}

		served = append(served, newCook)
	}
	return c.Render("main/viewTransaction", fiber.Map{"cooked": served})
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

	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	newTransaction := models.Transaction{
		IdUser: sess.Get("IDuser").(uint),
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

		temptot := DWislis.Quantity * productReal.ProductPrice
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

		productReal.Sold += DWislis.Quantity
		productReal.ProductStock -= DWislis.Quantity
		if err := initializers.GetDB().Save(&productReal).Error; err != nil {
			return err
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
