package handler

import (
	Response "Golang/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StructB struct {
	Name string `json:"name"`
}

func GetDataB(c *gin.Context) {
	var request StructB

	key := c.Param("key")
	c.Bind(&request)
	log.Println(request)

	c.JSON(http.StatusOK, Response.ReturnResponse(request.Name, key))
}
