package routes

import (
	"github.com/gin-gonic/gin"
	"loan-service/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/loans", controllers.GetLoans)
	router.POST("/loans", controllers.CreateLoan)
}
