package api

import (
	"blogging-platform/config"
	"blogging-platform/db"
	"blogging-platform/models"

	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the API routes and handlers
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define API routes
	router.GET("/posts", getPostsHandler)
	router.POST("/posts", createPostHandler)
	// Define more routes as needed

	return router
}



// Handler for creating a new post
func createPostHandler(c *gin.Context) {
    var newPost models.Post
    // Bind request body to Post struct
    if err := c.ShouldBindJSON(&newPost); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Connect to the database
    dbConn, err := db.ConnectDB(config.NewDBConfig())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
        return
    }
    defer dbConn.Close()

    // Prepare the SQL statement to insert the new post
    stmt := `
        INSERT INTO posts (title, author, content)
        VALUES (?, ?, ?)
    `
    // Execute the SQL statement
    if err := dbConn.Exec(stmt, newPost.Title, newPost.Author, newPost.Content).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
        return
    }

    // Return JSON response with the created post
    c.JSON(http.StatusCreated, newPost)
}

func getPostsHandler(c *gin.Context) {
    // Connect to the database
    dbConn, err := db.ConnectDB(config.NewDBConfig())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
        return
    }
    defer dbConn.Close()

    // Fetch all posts from the database
    var posts []models.Post
    if err := dbConn.Find(&posts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
        return
    }

    // Return JSON response with the fetched posts including IDs
    c.JSON(http.StatusOK, posts)
}


