package controllers

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/GoSpire/app/models"
	"github.com/uzzalhcse/GoSpire/bootstrap/app"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
)

type TestController struct {
	db *gorm.DB
}

func NewTestController() *TestController {
	return &TestController{
		db: app.DB,
	}
}

func (ctrl TestController) Test(c *gin.Context) {
	for i := 0; i < 2000; i++ {
		test := models.Post{
			Name:        faker.Name(),
			Slug:        faker.Username(),
			Description: faker.Paragraph(),
		}

		ctrl.db.Create(&test)
	}
	c.JSON(200, gin.H{
		"message": "Test Console",
	})
}
func (ctrl TestController) Upload(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	i := 0
	for {
		fileKey := fmt.Sprintf("files[%d][file]", i)
		priceKey := fmt.Sprintf("files[%d][price]", i)

		file, err := c.FormFile(fileKey)
		if err != nil {
			if err == http.ErrMissingFile {
				break // No more files, exit the loop
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error processing file %d: %s", i, err.Error())})
			return
		}

		price := c.PostForm(priceKey)

		// Handle the file
		if err := handleFile(file, price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		i++
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded"})
}
func handleFile(file *multipart.FileHeader, price string) error {
	// TODO File upload Logic
	fmt.Printf("File uploaded: %s, Price: %s\n", file.Filename, price)
	return nil
}
func (ctrl TestController) GetPosts(c *gin.Context) {
	var items []models.Post
	ctrl.db.Order("id DESC").Limit(500).Find(&items)
	c.JSON(200, gin.H{
		"message": "Test Success",
		"data":    items,
	})
}

func (ctrl TestController) BulkPostInsert(c *gin.Context) {
	for i := 0; i < 2000; i++ {
		test := models.Post{
			Name:        faker.Name(),
			Slug:        faker.Username(),
			Description: faker.Paragraph(),
		}

		ctrl.db.Create(&test)
	}
	var items []models.Post

	// Query the last 100 records from the database
	ctrl.db.Order("id DESC").Limit(500).Find(&items)

	c.JSON(200, gin.H{
		"message": "Test Success",
		"data":    items,
	})
}
