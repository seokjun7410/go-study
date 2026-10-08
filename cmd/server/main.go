package main

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/repository"
)

func main() {
	repository.InitDb()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":9100")
}
