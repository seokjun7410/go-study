package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()

	engine.GET("/pgin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Go Web Server!",
		})
	})

	engine.Run(":9100")
}
