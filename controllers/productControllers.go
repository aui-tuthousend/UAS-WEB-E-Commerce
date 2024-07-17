package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"web_uas/initializers"
	"web_uas/models"
)

func HashId(id string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(id), 14)
	return string(bytes), err
}

func ViewKategori(c *fiber.Ctx) error {
	idK := c.Params("idK")
	id, err := strconv.ParseUint(idK, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid category product ID")
	}

	var products []models.Product
	query := initializers.GetDB().Model(&models.Product{}).Where("category_id = ?", uint(id))
	if err := query.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("xxD")
	}

	count := len(products)

	var kategori models.Category
	if err := initializers.GetDB().First(&kategori, uint(id)).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Product not found")
	}

	//fmt.Print(idUser)

	return c.Render("main/viewKategori", fiber.Map{"products": products, "kategori": kategori, "len": count})
}

func ViewProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}

	var product models.Product
	if err := initializers.GetDB().First(&product, uint(id)).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Product not found")
	}

	var photos []models.PhotoProduct
	query := initializers.GetDB().Model(&models.PhotoProduct{}).Where("id_product = ?", product.ID)

	if err := query.Find(&photos).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("No photos found with the given conditions")
	}

	return c.Render("main/viewProduct", fiber.Map{"product": product, "photos": photos})
}

func StoreProduct(c *fiber.Ctx) error {
	const MaxFileSize = 10 * 1024 * 1024

	name := c.FormValue("name")
	desc := c.FormValue("desc")
	sto := c.FormValue("stok")
	pric := c.FormValue("price")
	cat := c.FormValue("category")

	stok, err := strconv.Atoi(sto)
	price, err := strconv.Atoi(pric)

	id, err := strconv.ParseUint(cat, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}

	file, err := c.FormFile("image")

	if file.Size > MaxFileSize {
		return c.Status(fiber.StatusRequestEntityTooLarge).SendString("File size exceeds the limit")
	}

	if err != nil {
		return err
	}
	imagePath := filepath.Join("images/cover", file.Filename)
	imagePath = strings.ReplaceAll(imagePath, "\\", "/")
	if err := c.SaveFile(file, imagePath); err != nil {
		return err
	}

	newProduct := models.Product{
		ProductName:        name,
		ProductDescription: desc,
		ProductImageCover:  imagePath,
		ProductStock:       stok,
		ProductPrice:       price,
		CategoryId:         uint(id),
		Sold:               0,
	}

	if err := initializers.GetDB().Create(&newProduct).Error; err != nil {
		return err
	}

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error retrieving multipart form:", err)
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse form")
	}
	files := form.File["images"]

	for _, file := range files {

		if file.Size > MaxFileSize {
			return c.Status(fiber.StatusRequestEntityTooLarge).SendString("File size exceeds the limit")
		}

		filePath := filepath.Join("images/productPhotos", file.Filename)
		filePath = strings.ReplaceAll(filePath, "\\", "/")
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file")
		}

		newPProduct := models.PhotoProduct{
			IdProduct: newProduct.ID,
			ImgPath:   filePath,
		}

		if err := initializers.GetDB().Create(&newPProduct).Error; err != nil {
			return err
		}
	}

	return c.Redirect("/")
}

func AddKategori(c *fiber.Ctx) error {
	kategori := c.FormValue("kategori")

	newKategori := models.Category{
		CategoryName: kategori,
	}

	if err := initializers.GetDB().Create(&newKategori).Error; err != nil {
		return err
	}
	return c.Redirect("/")
}
