package main

import (
	"gin-web-api/db"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin is working!",
		})
	})

	db.Init()

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
