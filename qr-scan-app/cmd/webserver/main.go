package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func main() {
	router := gin.Default()

	// Serve static files from the "assets" directory
	router.Static("/assets", "./assets")

	// Ruta para servir index.html como la interfaz principal
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Ruta para servir el formulario de generación de QR
	router.GET("/generador", func(c *gin.Context) {
		c.HTML(http.StatusOK, "generador.html", nil)
	})

	// Ruta para generar el código QR
	router.POST("/generate", func(c *gin.Context) {
		data := c.PostForm("data")
		qrCode, err := qrcode.Encode(data, qrcode.Medium, 256)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error generando el código QR")
			return
		}
		c.Header("Content-Type", "image/png")
		c.Writer.Write(qrCode)
	})

	router.Run(":8080")
}
