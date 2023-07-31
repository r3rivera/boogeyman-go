package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HealthcheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message":           "Running",
		"Private Key Found": fileExists("/var/local/private1_key.pem"),
		"Public Key Found":  fileExists("/var/local/public1_key.pem"),
	})
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
