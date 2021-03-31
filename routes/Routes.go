package routes

import (
	"dasargo/controllers"
	"github.com/gin-gonic/gin"
)

// SetUpRouter ...
func SetUpRouter() *gin.Engine {
	router := gin.Default();
	firstGroup := router.Group("/api") 
	{
		firstGroup.GET("user", controllers.GetUsers)
		firstGroup.POST("user", controllers.CreateUser)
		firstGroup.GET("user/:id", controllers.GetUserByID)
		firstGroup.PUT("user/:id", controllers.UpdateUser)
		firstGroup.DELETE("user/:id", controllers.DeleteUser)
	}
	return router
}