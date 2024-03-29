package resolvers

import (
	"context"
	"strings"

	"github.com/monopayments/iban.im/config"

	"github.com/monopayments/iban.im/handler"
	"github.com/monopayments/iban.im/model"
	// "fmt"
)

// ChangeProfile mutation change profile
func (r *Resolvers) ChangeProfile(ctx context.Context, args changeProfileMutationArgs) (*ChangeProfileResponse, error) {
	userID := ctx.Value(handler.ContextKey("UserID"))
	// fmt.Println("inside change profile")
	// fmt.Println("User id :",userID)
	// fmt.Printf("ctx: %+v\n",ctx)

	if userID == nil {
		msg := "Not Authorized"
		return &ChangeProfileResponse{Status: false, Msg: &msg, User: nil}, nil
	}
	user := model.User{}
	if err := config.DB.First(&user, userID).Error; err != nil {
		msg := "Not existing user"
		return &ChangeProfileResponse{Status: false, Msg: &msg, User: nil}, nil
	}

	if args.Bio != nil {
		user.Bio = *args.Bio
	}
	if args.Handle != nil {
		user.Handle = strings.ToLower(*args.Handle)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		msg := err.Error()
		return &ChangeProfileResponse{Status: false, Msg: &msg, User: nil}, err
	}
	return &ChangeProfileResponse{Status: true, Msg: nil, User: &UserResponse{u: &user}}, nil
}

type changeProfileMutationArgs struct {
	Bio    *string
	Handle *string
}

// ChangeProfileResponse is the response type
type ChangeProfileResponse struct {
	Status bool
	Msg    *string
	User   *UserResponse
}

// Ok for ChangeProfileResponse
func (r *ChangeProfileResponse) Ok() bool {
	return r.Status
}

// Error for ChangeProfileResponse
func (r *ChangeProfileResponse) Error() *string {
	return r.Msg
}
