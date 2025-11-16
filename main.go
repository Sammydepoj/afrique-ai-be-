package main

import (
	"fmt"

	"github.com/afrique-ai-be/db"
	"github.com/afrique-ai-be/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080

	fmt.Println("Server running on port :8080")
}
