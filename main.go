package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gin-tailwind-templ/templates"
)

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Render the home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", templates.HomePage())
	})

	// Start the server
	r.Run(":8080")
}
