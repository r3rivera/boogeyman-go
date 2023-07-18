package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r3rivera/boogeyman/services"
	"github.com/r3rivera/boogeyman/services/jwt"
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
	authUser(c, loginRequest.Email, loginRequest.Password)
}

func LoginBasicAuthHandler(c *gin.Context) {
	user, password, _ := c.Request.BasicAuth()
	if user == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Missing Credentials"})
		return
	}
	authUser(c, user, password)
}

func authUser(c *gin.Context, email, password string) {
	success, err := services.VerifyUserCredential(email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "System Error"})
		return
	}
	if !success {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Credentials"})
		return
	}
	if success {
		//Generate JWS token
		claims := map[string]string{
			"role": "user",
			"app":  "app1",
		}

		path := "/Users/r2devops/Devops/projects/golang/boogeyman-go/private_key.pem"
		jwtCert := jwt.NewCertFile(email, path, claims)
		jws, err := jwtCert.GenerateJWT()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "System Error"})
			return
		}
		//Get the default role and generate a token
		c.JSON(http.StatusOK, gin.H{
			"Message": "Success",
			"Token":   jws,
		})
		return
	}

}
