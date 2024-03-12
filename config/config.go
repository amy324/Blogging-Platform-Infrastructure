package config

// DBConfig holds the configuration settings for the MySQL database
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// NewDBConfig creates a new instance of DBConfig with default values
func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "password",
		DBName:   "blog_db",
	}
}


// APIConfig holds the configuration settings for the API service
type APIConfig struct {
	Port string
}

// LoadConfig loads the configuration settings for the API service
func LoadConfig() *APIConfig {
	// You can set default values here or load them from environment variables or a configuration file
	return &APIConfig{
		Port: "8080", // Default port 8080
	}
}
