package main

import (
	"os"

	"example.com/practice/initializers"
	"example.com/practice/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.GetRoute(r)

	r.Run(":" + os.Getenv("PORT"))
}
