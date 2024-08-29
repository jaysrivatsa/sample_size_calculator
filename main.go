package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaysrivatsa/sample_size_calculator/routes"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080
}
