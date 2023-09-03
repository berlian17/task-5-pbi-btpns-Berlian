package main

import (
	"task-5-pbi-btpns-Berlian/controllers"
	"task-5-pbi-btpns-Berlian/database"
	"task-5-pbi-btpns-Berlian/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.DBConnect()
	r := gin.Default()
	public := r.Group("/api")

	public.POST("register", controllers.Register)
	public.POST("login", controllers.Login)

	protected := r.Group("/api/users")
	protected.Use(middlewares.JWTAuth())
	protected.GET("/", controllers.GetCurrentUser)

	r.Run(":8080")
}
