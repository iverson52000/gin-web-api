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

type UserLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// var userModel = new(models.UserModel)

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

func Login(c *gin.Context) {
	var input UserLoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	res := db.GetDB().Where("email = ?", input.Email).First(&user)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
		return
	}

	bytePassword := []byte(input.Password)
	byteHashedPassword := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password"})
		return
	}

	// create and save token in redis
	// userFromDB, token, err := userModel.CreateUserToken(int64(user.ID))
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details"})
	// 	return
	// }

	// fmt.Println(userFromDB, token)

	c.JSON(http.StatusOK, gin.H{"message": "Login successiful"})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out!"})
}

func AllUser(c *gin.Context) {
	var users []models.User

	db.GetDB().Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
