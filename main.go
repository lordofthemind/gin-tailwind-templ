package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gin-tailwind-templ/templates" // Adjust the import path
)

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Render the home page
	r.GET("/", func(c *gin.Context) {
		// Use templ to render the template directly
		templates.HomePage().Render(c.Request.Context(), c.Writer)
	})

	// Start the server
	r.Run(":8080")
}
