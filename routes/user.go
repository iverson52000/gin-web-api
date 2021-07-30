package routes

import (
	"gin-web-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	v := r.Group("/v1")
	{
		v.POST("/user/register", controllers.Register)
		v.POST("/user/login", controllers.Login)
		v.POST("/user/logout", controllers.Logout)
		v.GET("/users", controllers.AllUser)
	}
}
