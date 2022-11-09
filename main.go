package main

import (
	Handler "Golang/Handler"
	"Golang/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(logger.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/getB/:key", Handler.GetDataB)

	r.Run(":9000")
}
