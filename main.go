package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"blogging-platform/api"
	"blogging-platform/config"
	"blogging-platform/db"
	"blogging-platform/monitoring"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	apiConfig := config.LoadConfig()
	dbConfig := config.NewDBConfig()

	fmt.Println("API Config:", apiConfig)
	fmt.Println("DB Config:", dbConfig)
	fmt.Println("")

	// Connect to the database
	dbConn, err := db.ConnectDB(dbConfig)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer dbConn.Close()

	fmt.Println("Connected to the database")
	fmt.Println("") // Print an empty line to flush the buffer

	// Setup metrics
	monitoring.SetupMetrics()

	// Create a new Gin router
	router := gin.New()

	// Define a handler for the root path
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Setup API routes and handlers
	api.SetupRouter(dbConn) // Pass dbConn to SetupRouter

	// Wrap the promhttp.Handler() with a custom handler function
	handler := func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	}

	// Register the custom handler with gin
	router.GET("/metrics", gin.HandlerFunc(handler))

	// Middleware to handle favicon.ico request
	router.Use(func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "favicon.ico") {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	})

	// Start HTTP server
	fmt.Println("Starting HTTP server on port", apiConfig.Port)
	err = http.ListenAndServe(":"+apiConfig.Port, router)
	if err != nil {
		log.Fatal("Error starting HTTP server:", err)
	}
}
