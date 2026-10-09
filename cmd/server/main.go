package main

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/repository"
	"go-api/internal/service"
)

func main() {
	repository.InitDb()

	boardRepo := repository.NewBoardRepository(repository.DB)
	boardService := service.NewBoardService(boardRepo)

	r := gin.Default()

	r.POST("/board/add", boardService.RegisterBoard)
	r.GET("/board/get/:id", boardService.GetBoardByID)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":9100")
}
