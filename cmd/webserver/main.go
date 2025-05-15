package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files from the "assets" directory
	router.Static("/assets", "./assets")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "QR Scan App",
		})
	})

	router.Run(":8080")
}

go get github.com/gin-gonic/gin