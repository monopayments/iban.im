package main // import "github.com/monocash/iban.im

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/monocash/iban.im/db"
	"github.com/monocash/iban.im/model"
	"github.com/monocash/iban.im/resolvers"
	"github.com/monocash/iban.im/schema"

	"github.com/appleboy/gin-jwt/v2"

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

	type login struct {
		Handle   string `form:"handle" json:"handle" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	var identityKey = "UserID"

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "ibanim zone",
		Key:         []byte("ibanim key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.Handle,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Handle: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user := model.User{}

			db.DB.Where("email = ?", loginVals.Handle).First(&user)

			if user.UserID == 0 {
				return "", jwt.ErrFailedAuthentication
			}

			if !user.ComparePassword(loginVals.Password) {
				return "", jwt.ErrFailedAuthentication
			}

			return &model.User{
				UserID:    user.UserID,
				LastName:  user.LastName,
				FirstName: user.FirstName,
			}, nil

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*model.User); ok && v.Admin == true {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

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
