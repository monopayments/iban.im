package main // import "github.com/monocash/iban.im

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/monocash/iban.im/db"
	"github.com/monocash/iban.im/resolvers"
	"github.com/monocash/iban.im/schema"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	context.Background()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router.GET("/graph", func(c *gin.Context) {
		c.HTML(http.StatusOK, "graph.tmpl.html", nil)
	})

	router.POST("/graph", func(c *gin.Context) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(c.Request.Body).Decode(&params); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
		schema := graphql.MustParseSchema(*schema.NewSchema(), &resolvers.Resolvers{DB: db}, opts...)

		response := schema.Exec(c, params.Query, params.OperationName, params.Variables)

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.JSON(200, response)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}

}
