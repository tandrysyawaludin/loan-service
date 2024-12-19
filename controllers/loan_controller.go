package controllers

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"loan-service/config"
	"loan-service/kafka"
	"loan-service/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var producer *kafka.Producer

// Initialize the producer once during application startup
func InitKafkaProducer(brokers []string) {
	var err error
	producer, err = kafka.NewProducer(brokers)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}
}

func ApproveLoan(c *gin.Context) {
	loanID := c.Param("id")

	// Perform business logic to approve the loan
	// For example, fetch the loan from DB, validate, update its state
	loan := models.GetLoanByID(loanID) // Mock function for getting loan data
	if loan == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Loan not found"})
		return
	}

	loan.State = "approved"
	// Save loan state to DB
	models.UpdateLoan(loan)

	// Publish an event to Kafka
	message := `{"loan_id": ` + loanID + `, "state": "approved"}`
	err := producer.SendMessage("loan-events", message)
	if err != nil {
		log.Printf("Failed to send Kafka message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to approve loan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan approved"})
}

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
