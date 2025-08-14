package main

import (
	"go-url-shrtr/db"
	"go-url-shrtr/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	defer server.Run(":8080")
}
