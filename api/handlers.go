package api

import (
	
	"blogging-platform/config"
	"blogging-platform/db"
	"blogging-platform/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRouter sets up the API routes and handlers
func SetupRouter(dbConn *gorm.DB) *gin.Engine {
	router := gin.Default()
	

	// Define API routes
	router.GET("/posts", func(c *gin.Context) {
		getPostsHandler(c, dbConn)
	})
	router.POST("/posts", func(c *gin.Context) {
		createPostHandler(c, dbConn)
	})
	router.PUT("/posts/:id", func(c *gin.Context) {
		updatePostHandler(c, dbConn)
	}) 
    router.DELETE("/posts/:id", func(c *gin.Context) {
		deletePostHandler(c, dbConn)
	})

	return router
}

// Handler for creating a new post
func createPostHandler(c *gin.Context, dbConn *gorm.DB) {
	var newPost models.Post
	// Bind request body to Post struct
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	// Return simplified response with the created post details
	c.JSON(http.StatusCreated, gin.H{
		"message": "Post Submitted",
		"title":   newPost.Title,
		"author":  newPost.Author,
		"content": newPost.Content,
	})
}

func getPostsHandler(c *gin.Context, dbConn *gorm.DB) {
	// Fetch all posts from the database
	var posts []models.Post
	if err := dbConn.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	// Return JSON response with the fetched posts including IDs
	c.JSON(http.StatusOK, posts)
}

func updatePostHandler(c *gin.Context, dbConn *gorm.DB) {
	// Get post ID from URL parameters
	postID := c.Param("id")

	// Fetch the post from the database
	var post models.Post
	if err := dbConn.Where("id = ?", postID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Bind request body to updated post struct
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the post in the database
	if err := dbConn.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	// Return JSON response with the updated post
	c.JSON(http.StatusOK, post)
}

// Handler for deleting a post
func deletePostHandler(c *gin.Context, dbConn *gorm.DB) {
	// Get post ID from URL parameters
	postID := c.Param("id")

	// Connect to the database
	// No need to defer closing connection as it's done in each handler
	dbConn, err := db.ConnectDB(config.NewDBConfig())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	// Fetch the post from the database
	var post models.Post
	if err := dbConn.Where("id = ?", postID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Delete the post from the database
	if err := dbConn.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	// Return JSON response indicating successful deletion
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}