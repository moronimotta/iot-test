package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Sample model for demonstration
type Device struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var db *gorm.DB

// Database connection function
func connectDB() {
	config := DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "root"),
		Database: getEnv("DB_NAME", "iot-schema"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully!")

	// Auto migrate the schema
	err = db.AutoMigrate(&Device{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

// Helper function to get environment variables with default values
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Health check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is running",
	})
}

// Get all devices
func getDevices(c *gin.Context) {
	var devices []Device
	if err := db.Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

// Create a new device
func createDevice(c *gin.Context) {
	var device Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&device).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, device)
}

// Get device by ID
func getDevice(c *gin.Context) {
	id := c.Param("id")
	var device Device

	if err := db.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}

	c.JSON(http.StatusOK, device)
}

// New endpoint to receive and display JSON data from Raspberry Pi
func receiveData(c *gin.Context) {
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON format",
			"details": err.Error(),
		})
		return
	}

	// Log the received data to console
	log.Printf("Received JSON data: %+v", jsonData)

	// Return the received data with additional metadata
	response := gin.H{
		"status":        "success",
		"message":       "Data received successfully",
		"timestamp":     fmt.Sprintf("%v", jsonData),
		"received_data": jsonData,
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	// Connect to database
	// connectDB()

	// Create Gin router
	r := gin.Default()

	// Health check route
	r.GET("/health", healthCheck)

	// New route to receive data from Raspberry Pi
	r.POST("/data", receiveData)

	// API routes
	api := r.Group("/api/v1")
	{
		api.GET("/devices", getDevices)
		api.POST("/devices", createDevice)
		api.GET("/devices/:id", getDevice)
	}

	// Start server on all interfaces
	port := getEnv("PORT", "8080")
	address := fmt.Sprintf("0.0.0.0:%s", port)

	log.Printf("Server starting on %s (listening on all interfaces)", address)
	log.Printf("Health check available at: http://0.0.0.0:%s/health", port)
	log.Printf("Data endpoint available at: http://0.0.0.0:%s/data", port)

	// Use Gin's Run method with the correct address to listen on all interfaces
	log.Fatal(r.Run(address))
}
