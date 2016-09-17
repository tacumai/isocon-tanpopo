package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "ok",
			"name":   name,
		})
	})

	router.Run()
}
