package main // import "github.com/monocash/iban.im

import (
	"context"
	"encoding/json"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/monocash/iban.im/config"
	"log"
	"net/http"

	"github.com/monocash/iban.im/handler"
	// "github.com/monocash/iban.im/model"
	"github.com/monocash/iban.im/resolvers"
	"github.com/monocash/iban.im/schema"

	"github.com/appleboy/gin-jwt/v2"

	"fmt"
	"github.com/gin-gonic/gin"
)
const identityKey = "UserID"


func main() {

	router := gin.New()

	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	defer config.DB.Close()

	context.Background()

	authMiddleware, err := handler.AuthMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.POST("/login", authMiddleware.LoginHandler)
	auth := router.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/profile", func(c *gin.Context) {
			c.HTML(http.StatusOK, "graph.tmpl.html", nil)
		})
	}

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func (c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			user, _ := c.Get(identityKey)

			c.JSON(200, gin.H{
				"userID":   claims[identityKey],
				"userName": user,
				"text":     "Hello World.",
			})
		})
	}

	router.GET("/graph", func(c *gin.Context) {
		fmt.Println("inside get graph")
		c.HTML(http.StatusOK, "graph.tmpl.html", nil)
	})
	
	authMW := authMiddleware.MiddlewareFunc()

	router.POST("/graph", func(c *gin.Context) {
		fmt.Println("inside post graph")
		ctx := c.Request.Context()

		if _, ok := c.Request.Header["Authorization"]; ok {
			authMW(c)
			
			fmt.Printf("c header auth: %+v\n",c.Request.Header.Get("Authorization"))
			claims := jwt.ExtractClaims(c)

			currentID,ok:=claims[identityKey].(float64)
			if !ok{
				currentID=0
			}
			ctx = context.WithValue(ctx,handler.ContextKey("UserID"), int(currentID))



		}
		
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(c.Request.Body).Decode(&params); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	
		opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
		schema := graphql.MustParseSchema(*schema.NewSchema(), &resolvers.Resolvers{}, opts...)

		response := schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.JSON(200, response)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",config.Config.App.Port),router))

}



