package main

import (
	"github.com/gin-gonic/gin"
	"org.gfoo/api-rest/db"
	"org.gfoo/api-rest/routes"
)

func main() {
	db.InitDB()
	defer db.Close()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
