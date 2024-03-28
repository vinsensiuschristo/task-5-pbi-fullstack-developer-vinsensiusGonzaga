package routes

import (
	"example.com/practice/controllers"
	"example.com/practice/middleware"
	"github.com/gin-gonic/gin"
)

func GetRoute(c *gin.Engine) {
	c.POST("/users/register", controllers.SignUp)
	c.POST("/users/login", controllers.Login)
	c.PUT("/users/:id", middleware.RequireAuth, middleware.AuthorizationMiddleware(), controllers.PutUser)
	c.DELETE("/users/:id", middleware.RequireAuth, middleware.AuthorizationMiddleware(), controllers.DeleteUser)
	c.GET("/photos", middleware.RequireAuth, controllers.GetPhoto)
	c.POST("/photos", middleware.RequireAuth, controllers.PostPhoto)
	c.PUT("/photos/:id", middleware.RequireAuth, controllers.PutPhoto)
	c.DELETE("/photos/:id", middleware.RequireAuth, middleware.AuthorizationMiddleware(), controllers.DeletePhoto)
}
