package main

import (
	"GO3/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Маршруты для регистрации и входа
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Защищенные маршруты с использованием JWT аутентификации
	protected := router.Group("/protected")
	protected.Use(handlers.AuthMiddleware())
	protected.GET("/items", handlers.GetItems) // Доступ только по токену
	protected.POST("/items", handlers.CreateItem)
	protected.PUT("/items/:id", handlers.UpdateItem)
	protected.DELETE("/items/:id", handlers.DeleteItem)

	router.Run(":3000")
}
