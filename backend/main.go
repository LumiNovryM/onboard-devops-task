package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

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

	// Check if environment variables are set
	checkEnv := utils.CHECK_ENV()
	if !checkEnv {
		utils.PRINT_LOG("ERROR", "FAILED CHECK ENV", utils.GetLocation(), "Failed To Load ENV, Please Check Configuration...")
	}

	connection.InitDB()

	gin.SetMode(gin.ReleaseMode)
	utils.PRINT_LOG("INFO", "SERVICE IS RUNNING", "", "")

	r := gin.Default()

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
