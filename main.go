package main // import "github.com/monocash/iban.im

import (
	"context"
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/monocash/iban.im/db"
	"github.com/monocash/iban.im/handler"
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

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(*schema.NewSchema(), &resolvers.Resolvers{DB: db}, opts...)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler.GraphiQL{})
	mux.Handle("/query", handler.Authenticate(&handler.GraphQL{Schema: schema}))

	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Println("Listening to... port " + port)
	if err = s.ListenAndServe(); err != nil {
		panic(err)
	}
	/*
		TODO: Use Gin
		router := gin.New()
		router.Use(gin.Logger())
		router.LoadHTMLGlob("templates/*.tmpl.html")
		router.Static("/static", "static")

		router.GET("/home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
		})
	*/

}
