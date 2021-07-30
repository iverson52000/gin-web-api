package routes

import (
	"gin-web-api/controllers"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(r *gin.Engine) {
	v := r.Group("/v1")
	{
		v.GET("/restaurants", controllers.AllRestaurants)
		v.GET("/restaurant/:id", controllers.OneRestaurant)
		v.POST("/restaurant", controllers.CreateRestaurant)
		v.PATCH("/restaurant/:id", controllers.UpdateRestaurant)
		v.DELETE("/restaurant/:id", controllers.DeleteRestaurant)
	}
}
