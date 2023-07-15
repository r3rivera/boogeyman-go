package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserReg struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var registration UserReg
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})

}
