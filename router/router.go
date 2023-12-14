package router

import "github.com/gin-gonic/gin"

func Inicialize() {
	// Inicializa o Router na config default do gin.
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}