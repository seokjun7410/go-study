package service

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-api/internal/model"
	"go-api/internal/repository"
	"time"
)

var jwtKey = []byte("my_secret_key_1234")

type AuthService struct {
	repo *repository.CompRepository
}

func NewAuthService(repo *repository.CompRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateToken(id string) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (s *AuthService) MakeToken(c *gin.Context) {
	var param model.Comp
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(500, gin.H{"code": "500", "message": err.Error()})
		return
	}

	if param.ID == "" {
		c.JSON(400, gin.H{"code": "500", "message": "Id is required"})
		return
	}

	comp, err := s.repo.FindByID(param.ID)
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "message": err.Error()})
		return
	}

	token, err := s.GenerateToken(comp.ID)
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "message": err.Error()})
		return
	}

	comp.Token = token
	if err := s.repo.Update(comp); err != nil {
		c.JSON(500, gin.H{"code": "500", "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": "200", "data": comp})
}
