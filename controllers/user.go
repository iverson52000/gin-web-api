package controllers

import (
	"gin-web-api/db"
	"gin-web-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	// Validate input
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	res := db.GetDB().Where("email = ?", input.Email).First(&user)
	if res.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	bytePassword := []byte(input.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong"})
		return
	}

	// Create user
	user = models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword), IsAdmin: false}
	db.GetDB().Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// func Login(c *gin.Context) {

// }

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out!"})
}

func AllUser(c *gin.Context) {
	var users []models.User

	db.GetDB().Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
