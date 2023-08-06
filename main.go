package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/r3rivera/boogeyman/controllers"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:*"
		},
		MaxAge: 12 * time.Hour,
	}))

	publicApi := r.Group("/api")
	// Health test
	publicApi.GET("/healthcheck", controllers.HealthcheckHandler)

	// Register User
	publicApi.POST("/register", controllers.RegisterUserHandler)

	// Login
	publicWeb := r.Group("/web")
	publicWeb.POST("/login", controllers.LoginUserHandler)
	publicWeb.POST("/login2", controllers.LoginBasicAuthHandler)

	//FileUpload
	publicWeb.PUT("/upload", controllers.HandleFileUploader)

	privateApi := publicApi.Group("/secured")
	privateApi.POST("/verify", controllers.VerifyJws)
	privateApi.POST("/profile", controllers.CreateUserProfile)
	privateApi.GET("/profile", controllers.GetUserDetails)

	return r
}

func main() {

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8082")
}
