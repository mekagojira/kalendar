package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Server = gin.Default()

func StartServer(port string) {
	Server.Use(cors.Default())
	startController()

	Server.Run(port)
}
