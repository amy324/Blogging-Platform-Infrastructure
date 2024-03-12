package db

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq" // Import PostgreSQL driver
    "blogging-platform/config"
    "blogging-platform/models"
)

// ConnectDB establishes a connection to the CockroachDB database using the provided configuration
func ConnectDB(config *config.DBConfig) (*gorm.DB, error) {
    connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", config.User, config.Password, config.Host, config.Port, config.DBName, config.SSLMode)
    fmt.Println("Database connection string:", connectionString)

    db, err := gorm.Open("postgres", connectionString)
    if err != nil {
        return nil, err
    }

    // Automatically create tables based on the defined models
    db.AutoMigrate(&models.Post{}) // Add more models here if needed

    fmt.Println("Connected to the database")
    return db, nil
}
