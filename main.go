package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Choose one of these:
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Golang!"})
	})
	// r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "Server is running!") })

	r.SetTrustedProxies(nil)
	r.Run(":8080")
}
