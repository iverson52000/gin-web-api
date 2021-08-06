package main

import (
	"fmt"
	"gin-web-api/db"
	"gin-web-api/middlewares"
	"gin-web-api/routes"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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

	// store := cookie.NewStore([]byte("secret"))
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("./public/html/*")
	r.Static("/public", "./public")

	routes.CommonRoutes(r)
	routes.RestaurantRoutes(r)
	routes.UserRoutes(r)

	r.Run(":" + port)

}
