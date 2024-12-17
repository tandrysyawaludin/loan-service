package controllers

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"loan-service/config"
	"loan-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLoans(c *gin.Context) {
	var loans []models.Loan
	if err := config.DB.Find(&loans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"request_id":  "1234",
			"status_code": 2,
			"message":     "Error fetching loans",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"request_id":  "1234",
		"status_code": 1,
		"message":     "success",
		"data":        loans,
	})
}

func CreateLoan(c *gin.Context) {
	var loan models.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"request_id":  "1234",
			"status_code": 2,
			"message":     "Invalid request payload",
		})
		return
	}

	loan.State = "proposed"
	if err := config.DB.Create(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"request_id":  "1234",
			"status_code": 2,
			"message":     "Error creating loan",
		})
		return
	}

	// Publish to Kafka
	msg, _ := json.Marshal(loan)
	_, _, err := config.KafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "loans",
		Value: sarama.StringEncoder(msg),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"request_id":  "1234",
			"status_code": 2,
			"message":     "Error publishing to Kafka",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"request_id":  "1234",
		"status_code": 1,
		"message":     "Loan created successfully",
		"data":        loan,
	})
}
