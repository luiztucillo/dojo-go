package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CryptoRequest struct {
	Value string `json:"value"`
}

type CryptoResponse struct {
	Value string `json:"value"`
}

func main() {
	app := gin.Default()

	app.POST("/encrypt", func(c *gin.Context) {
		var body CryptoRequest

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, &CryptoResponse{
			Value: Encrypt(body.Value),
		})
	})

	app.POST("/decrypt", func(c *gin.Context) {})

	app.Run(":3333")
}
