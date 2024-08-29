package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/additive", addtiveHandler)
	server.POST("/proportion", proportionHandler)
}
