package main

import (
	"go-jwt-authentication/controllers"
	"go-jwt-authentication/database"
	"go-jwt-authentication/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Iniitalize Database
	database.Connect("root:root@tcp(localhost:3306)/go_jwt_demo?parseTime=true") // should get value from config env
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
