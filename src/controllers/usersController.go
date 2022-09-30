package controllers

import (
	"github.com/gin-gonic/gin"

	database "github.com/lucca-rodrigues/myStock-API-GO/src/Database"
	"github.com/lucca-rodrigues/myStock-API-GO/src/models"
)


func GetAll(c *gin.Context){
	// db := database.GetDatabase()

	db := database.GetDBConnection()

	var users []models.User
	err := db.Find(&users).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

// func Create(c *gin.Context){
// 	db := database.GetDBConnection()
// 	var user models.User

// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "cannot bind JSON: " + err.Error(),
// 		})
// 		return
// 	}

// 	err = db.Create(&user).Error
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "cannot create book: " + err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, user)
// }