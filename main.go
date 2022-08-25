package main

import (
	"fmt"
	"go-jwt-authentication/controllers"
	"go-jwt-authentication/database"
	"go-jwt-authentication/middlewares"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file when MODE is not development|staging|sandbox|production
	log.Println("App Mode:", os.Getenv("MODE"))
	if os.Getenv("MODE") != "development" &&
		os.Getenv("MODE") != "staging" &&
		os.Getenv("MODE") != "sandbox" &&
		os.Getenv("MODE") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	// Iniitalize Database
	database.Connect(fmt.Sprintf(
		"root:root@tcp(%s:%s)/go_jwt_demo?parseTime=true",
		os.Getenv("MYSQL_PROVIDER_HOST"),
		os.Getenv("MYSQL_PROVIDER_PORT"),
	))
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
