package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gin-tailwind-templ/internals/templates"
)

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Render the home page
	r.GET("/", func(c *gin.Context) {
		templates.HomePage("Gin + Tailwind + Templ", "This is a dynamic message!").Render(c.Request.Context(), c.Writer)
	})

	r.GET("/api/data", func(c *gin.Context) {
		c.String(http.StatusOK, "HTMX response!")
	})

	// Start the server
	r.Run(":8080")
}
