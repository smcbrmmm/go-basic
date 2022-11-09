package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	type StructB struct {
		B           string `json:"b"`
		QueryString string `form:"query"`
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongping",
		})
	})

	r.POST("/getb/:query", func(c *gin.Context) {
		key := c.Param("query")
		var b StructB
		c.Bind(&b)

		c.JSON(200, gin.H{
			"b":     b.B,
			"query": key,
		})
	})

	r.Run()
}
