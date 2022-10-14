package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucca-rodrigues/myStock-API-GO/src/database"
	"github.com/lucca-rodrigues/myStock-API-GO/src/models"
)

func FindAllProducts(c *gin.Context){
	db := database.GetDatabase()

	var products []models.Product
	err := db.Find(&products).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, products)
}

func CreateProduct(c *gin.Context){
	db := database.GetDatabase()
	var product models.Product

	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&product).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, product)
}