package main

import (
	"backend-go/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Name interface {
	joji() string
}

func main() {
	log := logger.New()
	gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	log.Info("port")
	router.Run()
}
