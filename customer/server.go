package main

import "github.com/gin-gonic/gin"

//LoginRequest is the type that defines the import type.
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//LoginResponse is the type that defines the import type.
type LoginResonse struct {
	Session string `json:"session"`
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.POST("/Login", func(c *gin.Context) {
		var json LoginRequest
		if c.BindJSON(&json) == nil {
			c.JSON(200, gin.H{"status": json.Username})
		}
	})

	// By default it serves on :8080 unless a
	router.Run()
}
