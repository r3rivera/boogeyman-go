package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r3rivera/boogeyman/domain"
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

	user := domain.UserRegistration{
		FirstName: registration.FirstName,
		LastName:  registration.LastName,
		Email:     registration.Email,
	}

	err := user.Register()
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})

}
