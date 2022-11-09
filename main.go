package main

import (
	Handler "Golang/Handler"
	"Golang/logger"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(logger.Logger())

	r.GET("/ping", func(c *gin.Context) {
		log.Println("ping")
		c.JSON(200, gin.H{
			"message": "pongpingpangping",
		})
	})

	r.POST("/getB/:key", Handler.GetDataB)

	r.Run()
}
