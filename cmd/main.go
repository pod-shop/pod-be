package main

import (
	"net/http"
	"pod-be/pkg/modules/user"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func main() {
	InitDB()
	r := gin.Default()

	// Load the GraphiQL HTML template
	r.LoadHTMLGlob("templates/*")

	h := handler.New(&handler.Config{
		Schema:   &user.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Route for GraphQL POST requests
	r.POST("/graphql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	// Route to serve GraphiQL IDE on browser for development
	r.GET("/sandbox", func(c *gin.Context) {
		c.HTML(http.StatusOK, "graphiql.html", nil)
	})
	r.Run() // Listen and Serve on 0.0.0.0:8080 by default
}
