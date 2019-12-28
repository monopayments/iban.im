package main // import "github.com/monocash/iban.im

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	// "time"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/monocash/iban.im/db"
	// "github.com/monocash/iban.im/model"
	"github.com/monocash/iban.im/resolvers"
	"github.com/monocash/iban.im/schema"
	"github.com/monocash/iban.im/handler"

	"github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
	"fmt"
	"reflect"
)
var identityKey = "UserID"
// type login struct {
// 	Handle   string `form:"handle" json:"handle" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }
// var database *db.DB

func main() {

	router := gin.New()

	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	database, err := db.ConnectDB()
	fmt.Printf("db: %+v:",database)
	if err != nil {
		panic(err)
	}

	defer database.Close()

	context.Background()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	

	

	authMiddleware, err := handler.AuthMiddleware(database)


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
			fmt.Println("inside hello")
			fmt.Printf("claims: %+v\n",claims)
			fmt.Printf("user: %+v\n",user)
			fmt.Printf("gin context in hello : %+v\n",c)
			fmt.Println("identityKey: ", identityKey)
			fmt.Println("claims[identityKey]: ", claims[identityKey])


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
	// type ContextKey string
	authMW := authMiddleware.MiddlewareFunc()

	router.POST("/graph", func(c *gin.Context) {
		fmt.Println("inside post graph")
		// fmt.Printf("c body: %+v\n",c.Request.Body)
		ctx := c.Request.Context()

		if _, ok := c.Request.Header["Authorization"]; ok {
			authMW(c)
			
			fmt.Printf("c header auth: %+v\n",c.Request.Header.Get("Authorization"))
			claims := jwt.ExtractClaims(c)
			user, _ := c.Get(identityKey)
		
		// ctx := context.WithValue(c,ContextKey("UserID"), 1)
		currentID,ok:=claims[identityKey].(float64)
		if !ok{
			currentID=0
		}
		fmt.Printf("Current ID Type = %v\n", currentID) 
		ctx = context.WithValue(ctx,handler.ContextKey("UserID"), int(currentID))
		fmt.Printf("context: %+v\n",ctx)
		fmt.Printf("c: %+v\n",c)

		
		fmt.Printf("claims: %+v\n",claims)
		fmt.Printf("user: %+v\n",user)


		}
		
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(c.Request.Body).Decode(&params); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		// fmt.Printf("c body: %+v\n",c.Request.Body)
		// fmt.Printf("params: %+v\n",params)
		opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
		schema := graphql.MustParseSchema(*schema.NewSchema(), &resolvers.Resolvers{DB: database}, opts...)

		response := schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
		fmt.Printf("response: %+v",string(response.Data))
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

func getContextDetails(c context.Context){
	rv := reflect.ValueOf(c)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}

	if rv.Kind() == reflect.Struct {
		for i := 0; i < rv.NumField(); i++ {
			f := rv.Type().Field(i)

			if f.Name == "key" {
				fmt.Println("key: ", rv.Field(i))
			}
			if f.Name == "Context" {
				
				// this is just a repetition of the above, so you can make a recursive
				// function from it, or for loop, that stops when there are no more
				// contexts to be inspected.
				
				rv := rv.Field(i)
				for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
					rv = rv.Elem()
				}

				if rv.Kind() == reflect.Struct {
					for i := 0; i < rv.NumField(); i++ {
						f := rv.Type().Field(i)

						if f.Name == "key" {
							fmt.Println("key: ", rv.Field(i))
						}
						// ...
					}
				}
			}
			
		}
	}
}

