package routes

import (
	"gin-web-api/middlewares"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func CommonRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"gin-web-api": "v0.01",
			"goVersion":   runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	auth := r.Group("/auth")
	auth.Use(middlewares.CookieAuthMiddleware())
	{
		auth.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Everything is ok",
			})
		})
	}
}
