package main

import (
	"fmt"
	//"blogging-platform/api"
	"blogging-platform/config"
	"blogging-platform/db"
)

func main() {
	apiConfig := config.LoadConfig()
	dbConfig := config.NewDBConfig()

	fmt.Println("API Config:", apiConfig)
	fmt.Println("DB Config:", dbConfig)

	// Connect to the database
	db, err := db.ConnectDB(dbConfig)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Connected to the database")
}
