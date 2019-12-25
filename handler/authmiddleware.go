package handler

import(
	"time"
	"fmt"
	"github.com/monocash/iban.im/db"
	"github.com/monocash/iban.im/model"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	

)
var identityKey = "UserID"

type login struct {
	Handle   string `form:"handle" json:"handle" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// the jwt middleware
func AuthMiddleware(database *db.DB)(*jwt.GinJWTMiddleware, error){
return jwt.New(&jwt.GinJWTMiddleware{
	Realm:       "ibanim zone",
	Key:         []byte("ibanim key"),
	Timeout:     time.Hour,
	MaxRefresh:  time.Hour,
	IdentityKey: identityKey,
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		fmt.Println("inside payload func")
		if v, ok := data.(*model.User); ok {
			return jwt.MapClaims{
				identityKey: v.Handle,
			}
		}
		return jwt.MapClaims{}
	},
	IdentityHandler: func(c *gin.Context) interface{} {
		fmt.Println("inside identity handler")
		claims := jwt.ExtractClaims(c)
		fmt.Printf("c header auth: %+v\n",c.Request.Header.Get("Authorization"))
		fmt.Printf("claims: %+v\n",claims)
		fmt.Printf("claims identityKey: %+v\n",claims[identityKey])
		return &model.User{
			Handle: claims[identityKey].(string),
		}
	},
	Authenticator: func(c *gin.Context) (interface{}, error) {
		fmt.Println("inside Authenticator")
		var loginVals login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", fmt.Errorf("fatih 2: %v ", jwt.ErrMissingLoginValues) 
		}

		user := model.User{}

		database.DB.Where("email = ?", loginVals.Handle).First(&user)
		fmt.Printf("user: %+v\n",user)
		if user.UserID == 0 {
			return "", fmt.Errorf("fatih 2: %v ", jwt.ErrFailedAuthentication)
		}

		if !user.ComparePassword(loginVals.Password) {
			return "", fmt.Errorf("fatih 3: %v ", jwt.ErrFailedAuthentication)
		}

		return &model.User{
			UserID:    user.UserID,
			LastName:  user.LastName,
			FirstName: user.FirstName,
		}, nil

		return "", fmt.Errorf("fatih 4: %v ", jwt.ErrFailedAuthentication)
	},
	Authorizator: func(data interface{}, c *gin.Context) bool {
		fmt.Println("inside Authorizator")
		fmt.Printf("data: %+v\n",data)
		fmt.Printf("c : %+v\n",c)
		if _, ok := data.(*model.User);ok{
			return true
		}

		return false
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code ":    code,
			"message": message,
		})
	},
	TokenLookup:   "header: Authorization, query: token, cookie: jwt",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
})
}


