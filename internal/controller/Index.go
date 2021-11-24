package controller

import "github.com/gin-gonic/gin"

// Pong pong
func Pong(content *gin.Context) {
	content.JSON(200, gin.H{"message": "pong"})
}
