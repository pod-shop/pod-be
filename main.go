package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"pod-be/cmd/db"
	"pod-be/pkg/modules/user"
	"syscall"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	db.InitDB()
	r := gin.Default()

	// Load the GraphiQL HTML template
	// r.LoadHTMLGlob("/Users/macos/dev/pod-be/templates/*")

	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	executableDir := filepath.Dir(executable)

	// Construct the path to the templates directory relative to the executable
	templatesDir := filepath.Join(executableDir, "templates")

	// Load templates from the calculated directory
	r.LoadHTMLGlob(filepath.Join(templatesDir, "*"))
	mergedSchema := mergeSchemas(&user.Schema)

	h := handler.New(&handler.Config{
		Schema:   mergedSchema,
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

	es, _ := elasticsearch.NewDefaultClient()
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
	if gin.Mode() == gin.DebugMode {
		ginPort := ":8080"
		go func() {
			// Start the server in a goroutine
			if err := r.Run(ginPort); err != nil {
				log.Fatalf("Failed to start server: %v", err)
			}
		}()

	}
	// Handle OS signals to gracefully shut down the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server shutting down...")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server gracefully stopped.")

}

// Function to merge schemas
// Function to merge schemas
func mergeSchemas(schemas ...*graphql.Schema) *graphql.Schema {
	// Define combined query fields
	combinedQueryFields := graphql.Fields{}
	for _, schema := range schemas {
		queryType := (*schema).QueryType()
		for fieldName, field := range queryType.Fields() {
			combinedQueryFields[fieldName] = convertFieldDefinitionToField(field)
		}
	}

	// Create combined schema
	combinedSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Query",
			Fields: combinedQueryFields,
		}),
	})
	if err != nil {
		log.Fatalf("Error creating schema: %v", err)
	}

	return &combinedSchema
}

// Function to convert *graphql.FieldDefinition to *graphql.Field
func convertFieldDefinitionToField(field *graphql.FieldDefinition) *graphql.Field {
	args := make(graphql.FieldConfigArgument)
	for _, arg := range field.Args {
		args[arg.PrivateName] = &graphql.ArgumentConfig{
			Type:         arg.Type,
			DefaultValue: arg.DefaultValue,
			Description:  arg.Description(),
		}
	}
	return &graphql.Field{
		Name:              field.Name,
		Description:       field.Description,
		Type:              field.Type,
		Args:              args,
		Resolve:           field.Resolve,
		DeprecationReason: field.DeprecationReason,
	}
}
