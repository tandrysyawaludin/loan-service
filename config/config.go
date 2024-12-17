package config

import (
	"fmt"
	"github.com/IBM/sarama"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"loan-service/models"
	"log"
)

var DB *gorm.DB
var KafkaProducer sarama.SyncProducer

func InitDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/loan_service_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	DB = db
	fmt.Println("Database connected successfully")
}

func InitKafka() {
	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Error initializing Kafka producer: %v", err)
	}
	KafkaProducer = producer
	fmt.Println("Kafka connected successfully")
}

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.Loan{},
		&models.Investment{},
		&models.Employee{},
	)
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	fmt.Println("Database migration completed successfully")
}
