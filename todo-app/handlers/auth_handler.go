package handlers

import (
	"todo-app/models"
	"todo-app/services"
	"todo-app/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginReq models.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := services.AuthenticateUser(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Type)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})
}
