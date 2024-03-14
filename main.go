package main

import (
	"fmt"
	"blogging-platform/api"
	"blogging-platform/config"
	"blogging-platform/db"
)

func main() {
	apiConfig := config.LoadConfig()
	dbConfig := config.NewDBConfig()

	fmt.Println("API Config:", apiConfig)
	fmt.Println("DB Config:", dbConfig)
	fmt.Println("") // Print an empty line to flush the buffer

	// Connect to the database
	dbConn, err := db.ConnectDB(dbConfig)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer dbConn.Close()

	fmt.Println("Connected to the database")
	fmt.Println("") // Print an empty line to flush the buffer

	// Setup API routes and handlers
	router := api.SetupRouter(dbConn) // Pass dbConn to SetupRouter

	// Start HTTP server
	fmt.Println("Starting HTTP server on port", apiConfig.Port)
	err = router.Run(":" + apiConfig.Port)
	if err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}
