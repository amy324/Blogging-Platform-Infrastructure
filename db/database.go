package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"blogging-platform/config" 
)

// ConnectDB establishes a connection to the MySQL database using the provided configuration
func ConnectDB(config *config.DBConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
