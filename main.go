package main

import (
	"fmt"
	"gin-web-api/db"
	"gin-web-api/middlewares"
	"gin-web-api/models"
	"gin-web-api/routes"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := os.Getenv("PORT")
	fmt.Printf(" env vars: \n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n",
		port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	db.InitDB()
	db.InitRedis()

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	//testing route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin is working!",
		})
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"gin-web-api": "v0.01",
			"goVersion":   runtime.Version(),
		})
	})

	r.LoadHTMLGlob("./public/html/*")
	r.Static("/public", "./public")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	routes.RestaurantRoutes(r)
	routes.UserRoutes(r)

	var test models.AuthModel
	td, _ := test.CreateToken(1)
	fmt.Printf("%+v\n", td)

	r.Run(":" + port)

}
