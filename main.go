package main

import (
	"fmt"
	"gin-web-api/controllers"
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
	port := os.Getenv("PORT")
	fmt.Printf(" env vars: \n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n",
		port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	db.Init()

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Gin is working!",
			})
		})
		v1.GET("/restaurants", controllers.AllRestaurants)
		v1.GET("/restaurant/:id", controllers.OneRestaurant)
		v1.POST("/restaurant", controllers.CreateRestaurant)
		v1.PATCH("/restaurant/:id", controllers.UpdateRestaurant)
		v1.DELETE("/restaurant/:id", controllers.DeleteRestaurant)
	}

	r.Run(":" + port)

}
