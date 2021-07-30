package main

import (
	"fmt"
	"gin-web-api/db"
	"gin-web-api/routes"
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

	port := os.Getenv("PORT")
	fmt.Printf(" env vars: \n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n",
		port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	db.Init()

	r := gin.Default()

	//testing route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin is working!",
		})
	})

	routes.RestaurantRoutes(r)
	routes.UserRoutes(r)

	r.Run(":" + port)

}
