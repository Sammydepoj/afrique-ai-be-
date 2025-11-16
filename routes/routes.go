package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", serverCheck)
	server.POST("/news-letter", saveNewsLetterEmail)
}
