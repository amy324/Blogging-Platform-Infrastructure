package config

// PostConfig holds the configuration settings for a post
type PostConfig struct {
    // Define fields relevant to post configuration here
    Title  string
    Author string
    // Add more fields as needed
}

// NewPostConfig creates a new instance of PostConfig with default values
func NewPostConfig() *PostConfig {
    return &PostConfig{
        // Set default values for fields here
        Title:  "Default Title",
        Author: "Default Author",
        // Set default values for additional fields here
    }
}
