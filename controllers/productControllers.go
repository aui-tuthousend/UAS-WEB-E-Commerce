package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"path/filepath"
	"strconv"
	"strings"
	"web_uas/initializers"
	"web_uas/models"
)

func ShowProduct(c *fiber.Ctx) error {

	var products []models.Product

	if err := initializers.GetDB().Find(&products).Error; err != nil {
		return err
	}
	return c.Render("main/home", fiber.Map{"products": products})
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

	var photos models.PhotosProduct
	query := initializers.GetDB().Model(&models.PhotosProduct{}).Where("id_product = ?", uint(id))
	if err := query.Find(&photos).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("No photos found with the given conditions")
	}

	var imagePaths []string
	if err := json.Unmarshal(photos.ImagePaths, &imagePaths); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to unmarshal image paths")
	}

	return c.Render("produk/viewProduct", fiber.Map{
		"product": product,
		"photos":  imagePaths,
	})
}

func StoreProduct(c *fiber.Ctx) error {

	name := c.FormValue("name")
	desc := c.FormValue("desc")
	idSel := c.FormValue("idS")
	sto := c.FormValue("stok")
	pric := c.FormValue("price")

	idSeler, err := strconv.Atoi(idSel)
	stok, err := strconv.Atoi(sto)
	price, err := strconv.Atoi(pric)

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	imagePath := filepath.Join("images/cover", file.Filename)
	if err := c.SaveFile(file, imagePath); err != nil {
		return err
	}

	newProduct := models.Product{
		ProductName:        name,
		ProductDescription: desc,
		ProductImageCover:  imagePath,
		IdSeller:           idSeler,
		ProductStock:       stok,
		ProductPrice:       price,
	}

	if err := initializers.GetDB().Create(&newProduct).Error; err != nil {
		return err
	}

	form, err := c.MultipartForm()
	files := form.File["images"]
	var imagePathsPhotos []string

	for _, file := range files {
		filePath := filepath.Join("images/productPhotos", file.Filename)
		imagePath = strings.ReplaceAll(imagePath, "\\", "/")
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file")
		}
		imagePathsPhotos = append(imagePathsPhotos, filePath)
	}

	imagePathsJSON, err := json.Marshal(imagePathsPhotos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to marshal image paths")
	}

	newPhotos := models.PhotosProduct{
		IdProduct:  newProduct.ID,
		ImagePaths: datatypes.JSON(imagePathsJSON),
	}
	if err := initializers.GetDB().Create(&newPhotos).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create post")
	}

	return c.Redirect("/")
}