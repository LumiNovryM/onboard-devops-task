package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/gin-contrib/cors"

	// Utils
	"backend/app/utils"
	"backend/connection"
	"backend/app/controllers"
)

func main() {

	// Print Backend Status 
	utils.PRINT_LOG("INFO", "Backend Service Start", utils.GetLocation(), "");

	// Check if environment variables are set on config.json
	if _, err := os.Stat("config.json"); err == nil {
		err := godotenv.Load()
		if err != nil {
			utils.PRINT_LOG("ERROR", "FAILED LOAD ENV", utils.GetLocation(), fmt.Sprintf(err.Error()))
		}
	}

	connection.InitDB()

	gin.SetMode(gin.ReleaseMode)
	utils.PRINT_LOG("INFO", "SERVICE IS RUNNING", "", "")

	r := gin.Default()

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Refuse Trace Method
	r.Use(utils.DISABLE_TRACE_METHOD())

	// Messages API
	// Get
	r.GET("/api/messages", controllers.GetAllMessages)
	// Post
	r.POST("/api/messages", controllers.CreateMessage)

	
	// Status API
	// Get
	r.GET("/api/status", controllers.GetAllStatus)

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

}
