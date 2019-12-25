package resolvers

import (
	"context"

	"github.com/monocash/iban.im/handler"
	"github.com/monocash/iban.im/model"
	"fmt"
	"reflect"
)

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
						}else{
							fmt.Printf("value: %+v\n", rv.Field(i))
						}
						// ...
					}
				}
			}
		}
	}
}

// GetMyProfile resolver
func (r *Resolvers) GetMyProfile(ctx context.Context) (*GetMyProfileResponse, error) {
	UserID := ctx.Value(handler.ContextKey("UserID"))
	jwtToken := ctx.Value(handler.ContextKey("JWT_TOKEN"))
	fmt.Println("inside resolver getmyprofile")
	fmt.Println("UserID: ",UserID)
	fmt.Println("JWT_TOKEN: ",jwtToken)
	fmt.Printf("context:%+v\n",ctx)
	// fmt.Printf("context Req :%+v\n",ctx.Request)
	getContextDetails(ctx)
	if UserID == nil {
		msg := "Not Authorized"
		return &GetMyProfileResponse{Status: false, Msg: &msg, User: nil}, nil
	}

	user := model.User{}
	if err := r.DB.First(&user, UserID).Error; err != nil {
		msg := "Not found"
		return &GetMyProfileResponse{Status: false, Msg: &msg, User: nil}, nil
	}
	return &GetMyProfileResponse{Status: true, Msg: nil, User: &UserResponse{u: &user}}, nil
}

// GetMyProfileResponse is the response type
type GetMyProfileResponse struct {
	Status bool
	Msg    *string
	User   *UserResponse
}

// Ok for GetMyProfileResponse
func (r *GetMyProfileResponse) Ok() bool {
	return r.Status
}

// Error for GetMyProfileResponse
func (r *GetMyProfileResponse) Error() *string {
	return r.Msg
}
