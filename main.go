package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Projeto gym em loading": "Contributers: Emanuel and Mauricio",
		})
	})

	r.Run(":8080")
}
