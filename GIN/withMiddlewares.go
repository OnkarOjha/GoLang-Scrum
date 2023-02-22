package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logger middleware that logs each request method and URL
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before request
		fmt.Printf("[%s] %s\n", c.Request.Method, c.Request.URL.Path)

		// Call the next middleware (or handler)
		c.Next()

		// After request
		status := c.Writer.Status()
		fmt.Printf("[%d] %s\n", status, http.StatusText(status))
	}
}

// Authentication middleware that checks for a specific header
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Call the next middleware (or handler)
		c.Next()
	}
}

// Example handler that returns a greeting
func Greet(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

func main() {
	// Create a new router with the Logger middleware
	router := gin.New()
	router.Use(Logger())

	// Set up a route that requires authentication
	auth := router.Group("/", Auth())
	auth.GET("/greet", Greet)

	// Start the server
	router.Run(":8000")
}
