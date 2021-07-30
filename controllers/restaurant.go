package controllers

import (
	"gin-web-api/db"
	"gin-web-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRestaurantInput struct {
	Name   string `json:"name" binding:"required"`
	Type   string `json:"type" binding:"required"`
	Rating int    `json:"rating" binding:"required"`
}

type UpdateRestaurantInput struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Rating int    `json:"rating"`
}

func AllRestaurants(c *gin.Context) {
	var restaurants []models.Restaurant

	db.GetDB().Find(&restaurants)

	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}

func OneRestaurant(c *gin.Context) {
	var restaurant models.Restaurant

	if err := db.GetDB().Where("id = ?", c.Param("id")).First(&restaurant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Restaurant not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}

func CreateRestaurant(c *gin.Context) {
	// Validate input
	var input CreateRestaurantInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create restaurant
	restaurant := models.Restaurant{Name: input.Name, Type: input.Type, Rating: input.Rating}
	db.GetDB().Create(&restaurant)

	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}

func UpdateRestaurant(c *gin.Context) {
	var restaurant models.Restaurant

	if err := db.GetDB().Where("id = ?", c.Param("id")).First(&restaurant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Restaurant not found!"})
		return
	}

	var input UpdateRestaurantInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update restaurant
	updatedRestaurant := models.Restaurant{Name: input.Name, Type: input.Type, Rating: input.Rating}
	db.GetDB().Model(&restaurant).Updates(updatedRestaurant)

	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}

func DeleteRestaurant(c *gin.Context) {
	// Validate restaurant
	var restaurant models.Restaurant

	if err := db.GetDB().Where("id = ?", c.Param("id")).First(&restaurant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Restaurant not found!"})
		return
	}

	db.GetDB().Delete(&restaurant)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
