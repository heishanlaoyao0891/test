package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/helloworld", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello, Gin!",
			"status":  "success",
		})
	})

	server.Run(":8081")
}
