package main

import "github.com/gin-gonic/gin"

//RegisterRequest is the type that defines the import type.
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.POST("/Register", func(c *gin.Context) {
		var json RegisterRequest
		if c.BindJSON(&json) == nil {
			c.JSON(200, gin.H{"status": json.Username})
		}
	})

	// By default it serves on :8080 unless a
	router.Run(":8081")
}
