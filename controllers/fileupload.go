package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FileMetaData struct {
	Description string `form:"description"`
}

func HandleFileUploader(c *gin.Context) {

	var fileMetaData FileMetaData
	err := c.ShouldBind(&fileMetaData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request in the file"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error getting the file"})
		return
	}

	filePath := "./uploads/" + file.Filename
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error processing the file"})
		return
	}
	defer os.Remove(filePath)

}
