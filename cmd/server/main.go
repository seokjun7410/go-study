package main

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/repository"
	"go-api/internal/service"
	"go-api/pkg/util"
)

func main() {
	util.InitZapLogger()
	repository.InitDb()

	boardRepo := repository.NewBoardRepository(repository.DB)
	boardService := service.NewBoardService(boardRepo)
	compRepo := repository.NewCompRepository(repository.DB)
	authService := service.NewAuthService(compRepo)

	r := gin.Default()
	authMiddleware := authService.AuthMiddleware()
	r.Group("/", authMiddleware)

	r.POST("/board/add", boardService.RegisterBoard)
	r.GET("/board/get/:id", boardService.GetBoardByID)
	r.POST("/auth/token", authService.MakeToken)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":9100")
}
