package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r3rivera/boogeyman/controllers"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	publicApi := r.Group("/api")
	// Health test
	publicApi.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Running")
	})

	// Register User
	publicApi.POST("/register", controllers.RegisterUserHandler)

	publicWeb := r.Group("/web")
	publicWeb.POST("/login", controllers.LoginUserHandler)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
