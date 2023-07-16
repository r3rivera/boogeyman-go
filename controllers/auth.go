package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r3rivera/boogeyman/services"
)

type LoginRequest struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func LoginUserHandler(c *gin.Context) {
	var loginRequest LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Missing Credentials"})
		return
	}

	success, err := services.VerifyUserCredential(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "System Error"})
		return
	}
	if !success {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Credentials"})
		return
	}

	if success {
		c.JSON(http.StatusOK, gin.H{"Message": "Success"})
		return
	}
}
