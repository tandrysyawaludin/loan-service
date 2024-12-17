package main

import (
	"github.com/gin-gonic/gin"
	"loan-service/config"
	"loan-service/routes"
)

func main() {
	// Initialize DB
	config.InitDB()

	// Run migrations
	config.MigrateDB()

	// Initialize Kafka
	config.InitKafka()

	// Setup Router
	router := gin.Default()
	routes.RegisterRoutes(router)

	// Start Server
	router.Run(":8080")
}
