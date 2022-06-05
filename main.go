package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gtokmak/jwt-rest-api/controllers"
	"github.com/gtokmak/jwt-rest-api/middlewares"
	"github.com/gtokmak/jwt-rest-api/models"
)

func main() {

	models.ConnectDataBase()

	gin.ForceConsoleColor()
	router := gin.Default()

	public := router.Group("/api")

	public.GET("/ping", controllers.Ping)
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	router.Run(":7777")

}
